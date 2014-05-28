package models

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
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

func InitDb(name string, user string, passwd string, logPrefix string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s", user, name, passwd))
	if err != nil {
		return nil, err
	}

	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	if err = dbMap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}

	dbMap.TraceOn("[sql]", log.New(os.Stdout, logPrefix+":", log.Lmicroseconds))
	return dbMap, nil
}
