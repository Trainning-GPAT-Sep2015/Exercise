package query

import (
	"time"

	r "github.com/dancannon/gorethink"
)

const (
	DB_NAME = "simple_blog"
)

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
