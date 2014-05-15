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

type Post struct {
	Id      int64  `json:"id",binding:"-"`
	Content string `json:"content",binding:"required"`
}

func (post *Post) Validate(errors *binding.Errors, req *http.Request) {
	if post.Content == "" {
		errors.Fields["content"] = "Content is missing"
	}
	if len(post.Content) > 140 {
		errors.Fields["content"] = "Content must be max 140 characters"
	}
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/tweets.db")
	checkErr(err)

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbMap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	err = dbMap.CreateTablesIfNotExists()
	checkErr(err)

	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	return dbMap
}
