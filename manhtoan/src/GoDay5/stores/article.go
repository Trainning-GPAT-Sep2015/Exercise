package stores

import (
	"GoDay5/api/rethinkdb"
)

type ArticleStore struct {
	re *rethinkdb.Instance
}

func NewArticleStore(re *rethinkdb.Instance) *ArticleStore {
	return &ArticleStore{
		re: re,
	}
}
