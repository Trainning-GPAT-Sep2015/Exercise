package querry

import (
	r "github.com/dancannon/gorethink"
	"time"
)

const (
	DB_NAME       = "blog"
	ARTICLE_TABLE = "article"
	TAG_TABLE     = "tag"
	AUTHOR_TABLE  = "author"
)

type Author struct {
	Name string `gorethink:"name"`
	Id   string `gorethink:"id"`
}

type Article struct {
	Title   string   `gorethink:"title"`
	Content string   `gorethink:"content"`
	Author  string   `gorethink:"author"`
	Tag     []string `gorethink:"tag"`
	Id      string   `gorethink:"id"`
}

func GetAllArticle(session *r.Session) ([]*Article, error) {
	cursor, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).Run(session)
	if err != nil {
		return nil, err
	}
	var articles []*Article
	cursor.All(&articles)
	return articles, nil
}

func GetAllAuthor(session *r.Session) ([]*Author, error) {
	cursor, err := r.DB(DB_NAME).Table(AUTHOR_TABLE).Run(session)
	if err != nil {
		return nil, err
	}
	var authors []*Author
	cursor.All(&authors)
	return authors, nil
}

func GetAllArticleByAuthor(author_id string, session *r.Session) ([]*Article, error) {
	cursor, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).GetAllByIndex("author_id", author_id).Run(session)
	if err != nil {
		return nil, err
	}
	var articles []*Article
	cursor.All(&articles)
	return articles, nil
}

func GetAllArticleByTag(tag string, session *r.Session) ([]*Article, error) {
	cursor, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).GetAllByIndex("tag", tag).Run(session)
	if err != nil {
		return nil, err
	}
	var articles []*Article
	cursor.All(&articles)
	return articles, nil
}

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
