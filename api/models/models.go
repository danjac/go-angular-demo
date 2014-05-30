package models

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const maxContentLength = 140

var dbMap *gorp.DbMap

type Post struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}

type ValidationResult struct {
	Errors map[string]string `json:"errors"`
	OK     bool              `json:"ok"`
}

func NewValidationResult() *ValidationResult {
	return &ValidationResult{make(map[string]string), true}
}

func (result *ValidationResult) Error(field string, msg string) {
	result.Errors[field] = msg
	result.OK = false
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

func (post *Post) Validate() *ValidationResult {

	result := NewValidationResult()

	if post.Content == "" {
		result.Error("content", "Content is missing")
	}
	if len(post.Content) > maxContentLength {
		result.Error("content", fmt.Sprintf("Content must be max %d characters", maxContentLength))
	}

	return result
}

func Init(db *sql.DB, logPrefix string) (*gorp.DbMap, error) {

	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	if err := dbMap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}
	dbMap.TraceOn("[sql]", log.New(os.Stdout, logPrefix+":", log.Lmicroseconds))
	return dbMap, nil

}
