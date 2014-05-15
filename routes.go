package main

import (
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

func getTweets(db gorp.SqlExecutor, r render.Render) {
	var posts []Post
	_, err := db.Select(&posts, "SELECT * FROM posts ORDER BY id DESC")
	checkErr(err)
	r.JSON(http.StatusOK, posts)
}

func addTweet(
	db gorp.SqlExecutor,
	post Post,
	errors binding.Errors,
	r render.Render) {
	if errors.Count() > 0 {
		r.JSON(http.StatusConflict, errors)
		return
	}
	err := db.Insert(&post)
	checkErr(err)
	r.JSON(http.StatusOK, post)
}

func deleteTweet(
	db gorp.SqlExecutor,
	params martini.Params,
	r render.Render) {
	obj, err := db.Get(Post{}, params["id"])
	checkErr(err)
	if obj == nil {
		r.JSON(http.StatusNotFound, "NotFound")
		return
	}
	_, err = db.Delete(obj)
	checkErr(err)
	r.JSON(http.StatusOK, "Deleted")
}

func addRoutes(m *martini.ClassicMartini) {

	m.Group("/api", func(r martini.Router) {
		r.Get("", getTweets)
		r.Post("", binding.Json(Post{}), addTweet)
		r.Delete("/:id", deleteTweet)
	})
}
