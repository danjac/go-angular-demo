package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
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

func initDb() (*gorp.DbMap, error) {
	db, err := sql.Open("sqlite3", "/tmp/tweets.db")
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbMap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}
    dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds)) 
	return dbMap, nil
}

func getTweets(dbMap *gorp.DbMap, r render.Render) {
	var posts []Post
	_, err := dbMap.Select(&posts, "SELECT * FROM posts ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}
	r.JSON(http.StatusOK, posts)
}

func addTweet(
	dbMap *gorp.DbMap,
	post Post,
	errors binding.Errors,
	r render.Render) {
	if errors.Count() > 0 {
		r.JSON(http.StatusConflict, errors)
		return
	}
	err := dbMap.Insert(&post)
	if err != nil {
		log.Fatal(err)
	}
	r.JSON(http.StatusOK, post)
}

func deleteTweet(
	dbMap *gorp.DbMap,
	params martini.Params,
	r render.Render) {
    obj, err := dbMap.Get(Post{}, params["id"])
	if err != nil {
		log.Fatal(err)
	}
    if obj == nil {
        r.JSON(http.StatusNotFound, "NotFound")
        return
    }
    _, err = dbMap.Delete(obj)
	if err != nil {
		log.Fatal(err)
	}
	r.JSON(http.StatusOK, "Deleted")
}


func addHandlers(m *martini.ClassicMartini) {

	m.Group("/api", func(r martini.Router) {
		r.Get("", getTweets)
		r.Post("", binding.Json(Post{}), addTweet)
		r.Delete("/:id", deleteTweet)
	})
}

func main() {

	// DATABASE
	dbMap, err := initDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dbMap.Db.Close()

	// SERVER
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Map(dbMap)
	addHandlers(m)
	m.Run()
}
