package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/martini-contrib/binding"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

var dbMap *gorp.DbMap

type Post struct {
	Id      int64  `json:"id",binding:"-"`
	Content string `json:"content",binding:"required"`
}

func GetPosts() ([]Post, error) {
	var posts []Post
	_, err := dbMap.Select(&posts, "SELECT * FROM posts ORDER BY id DESC")
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

func (post *Post) Validate(errors *binding.Errors, req *http.Request) {
	if post.Content == "" {
		errors.Fields["content"] = "Content is missing"
	}
	if len(post.Content) > 140 {
		errors.Fields["content"] = "Content must be max 140 characters"
	}
}

func initDb(dbName string) *gorp.DbMap {
	db, err := sql.Open("sqlite3", dbName)
	checkErr(err)

	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbMap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	err = dbMap.CreateTablesIfNotExists()
	checkErr(err)

	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	return dbMap
}
