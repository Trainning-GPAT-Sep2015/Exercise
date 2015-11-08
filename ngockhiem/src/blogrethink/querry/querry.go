package querry

import (
	"fmt"
	r "github.com/dancannon/gorethink"
)

const (
	DB_NAME       = "blog"
	ARTICLE_TABLE = "article"
	TAG_TABLE     = "tag"
	AUTHOR_TABLE  = "author"
)

type Author struct {
	name     string
	id       float64
	articles []float64
}

type Article struct {
	title   string
	content string
	author  float64
	tags    []float64
	id      float64
}

type Tag struct {
	name     string
	articles []float64
}

func GetAllArticle(session *r.Session) ([]Article, error) {
	res, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).Run(session)
	if err != nil {
		return []Article{}, err
	}
	defer res.Close()
	var articles []Article
	var row interface{}
	for res.Next(&row) {
		var article Article
		if convert, ok := row.(map[string]interface{}); ok {
			article.title = convert["title"].(string)
			article.content = convert["content"].(string)
			article.id = convert["id"].(float64)
			if convert["author"] != nil {
				article.author = convert["author"].(float64)
			}
			if convert["tags"] != nil {
				tags := convert["tags"].([]interface{})
				for _, tag := range tags {
					article.tags = append(article.tags, tag.(float64))
				}
			}
			articles = append(articles, article)
		} else {
			fmt.Println("Cannot convert article row")
		}
	}
	return articles, nil
}
func GetAllAuthor(session *r.Session) ([]Author, error) {
	res, err := r.DB(DB_NAME).Table(AUTHOR_TABLE).Run(session)
	if err != nil {
		return []Author{}, err
	}
	defer res.Close()
	var authors []Author
	var row interface{}
	for res.Next(&row) {
		var author Author
		if convert, ok := row.(map[string]interface{}); ok {
			author.id = convert["id"].(float64)
			author.name = convert["name"].(string)
			if convert["articles"] != nil {
				articles := convert["articles"].([]interface{})
				for _, article := range articles {
					author.articles = append(author.articles, article.(float64))
				}
			}
			authors = append(authors, author)
		} else {
			fmt.Println("Cannot convert row")
		}
	}
	res.All(&authors)
	return authors, nil
}

func GetAuthorArticle(Author author) ([]Article, error) {
	res, err := r.DB(DB_NAME).Table(ARTICLE_TABLE).Run(session)
	if err != nil {
		return []Article{}, err
	}
	defer res.Close()
	var articles []Article
	var row interface{}
	for res.Next(&row) {
		var article Article
		if convert, ok := row.(map[string]interface{}); ok {
			article.title = convert["title"].(string)
			article.content = convert["content"].(string)
			article.id = convert["id"].(float64)
			if convert["author"] != nil {
				article.author = convert["author"].(float64)
			}
			if convert["tags"] != nil {
				tags := convert["tags"].([]interface{})
				for _, tag := range tags {
					article.tags = append(article.tags, tag.(float64))
				}
			}
			articles = append(articles, article)
		} else {
			fmt.Println("Cannot convert article row")
		}
	}
	return articles, nil
}
