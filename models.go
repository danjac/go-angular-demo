package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var dbMap *gorp.DbMap

type Post struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}

type Errors struct {
	Fields map[string]string
}

func NewErrors() *Errors {
	return &Errors{Fields: make(map[string]string)}
}

func (errors Errors) Count() int {
	return len(errors.Fields)
}

func GetPosts() ([]Post, error) {
	var posts []Post
	_, err := dbMap.Select(&posts, "SELECT * FROM posts ORDER BY id DESC LIMIT 20")
	return posts, err
}

func GetPost(postId string) (*Post, error) {
	obj, err := dbMap.Get(Post{}, postId)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, nil
	}
	return obj.(*Post), nil
}

func (post *Post) Save() error {
	return dbMap.Insert(post)
}

func (post *Post) Delete() error {
	_, err := dbMap.Delete(post)
	return err
}

func (post *Post) Validate() *Errors {

	errors := NewErrors()

	if post.Content == "" {
		errors.Fields["content"] = "Content is missing"
	}
	if len(post.Content) > 140 {
		errors.Fields["content"] = "Content must be max 140 characters"
	}

	return errors
}

func InitDb(dbName string) (*gorp.DbMap, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbMap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	if err = dbMap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}

	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	return dbMap, nil
}
