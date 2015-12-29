package Query

import r "github.com/dancannon/gorethink"
import (
	"time"
)

const (
	DB_NAME        = "blog"
	ARTICLE_TABLE  = "article"
	AUTHOR_TABLE   = "author"
	INDEX_AUTHORID = "author_id"
	INDEX_TAG      = "tag"
)

type Author struct {
	Name string `gorethink:"name"`
	Id   string `gorethink:"id"`
}

type Article struct {
	Title    string   `gorethink:"title"`
	Content  string   `gorethink:"content"`
	AuthorId string   `gorethink:"author_id"`
	Tags     []string `gorethink:"tags"`
	Id       string   `gorethink:"id"`
}

func GetAllArticles(session *r.Session) (articles []*Article, err error) {
	cursor, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).Run(session)
	if err != nil {
		return nil, err
	}

	err = cursor.All(&articles)
	return
}

func GetAllAuthors(session *r.Session) (authors []*Author, err error) {
	cursor, err := r.DB(DB_NAME).Table(AUTHOR_TABLE).Run(session)
	if err != nil {
		return nil, err
	}

	err = cursor.All(&authors)
	return
}

func GetArticlesByAuthor(session *r.Session, authorId string) (articles []*Article, err error) {
	cursor, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).GetAllByIndex(INDEX_AUTHORID, authorId).Run(session)
	if err != nil {
		return nil, err
	}

	err = cursor.All(&articles)
	return
}

func GetArticlesByTag(session *r.Session, tag string) (articles []*Article, err error) {
	cursor, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).GetAllByIndex(INDEX_TAG, tag).Run(session)
	if err != nil {
		return nil, err
	}

	err = cursor.All(&articles)
	return
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
type Message struct {
	Content     string    `gorethink:"content"`
	CreatedTime time.Time `gorethink:"created_time"`
}

func GetMessagesByRoom(session *r.Session, roomId int) (messages []*Message, err error) {
	cursor, err := r.DB(DB_NAME).Table("message").
		Between([]interface{}{roomId, r.MinVal}, []interface{}{roomId, r.MaxVal}, r.BetweenOpts{
		Index: "room_time",
	}).
		OrderBy(r.OrderByOpts{Index: "room_time"}).
		Limit(30).
		Run(session)

	if err != nil {
		return nil, err
	}

	err = cursor.All(&messages)
	return
}

func SubscribeMessages(session *r.Session, roomId int, callback func(err error, message *Message)) (*r.Cursor, error) {
	cursor, err := r.DB(DB_NAME).Table("message").
		Between([]interface{}{roomId, r.MinVal}, []interface{}{roomId, r.MaxVal}, r.BetweenOpts{
		Index: "room_time",
	}).
		Changes().
		Run(session)

	if err != nil {
		return nil, err
	}

	go func() {
		var change struct {
			NewVal *Message `gorethink:"new_val"`
			OldVal *Message `gorethink:"old_val"`
		}
		for {
			ok := cursor.Next(&change)
			if !ok {
				err := cursor.Err()
				cursor.Close()
				callback(err, nil)
				return
			}

			if change.NewVal != nil {
				callback(nil, change.NewVal)
			}
		}
	}()

	return cursor, nil
}
