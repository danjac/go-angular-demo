package main

import (
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

func getTweets(db gorp.SqlExecutor, r render.Render) {
    posts, err := GetPosts(db)
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
	err := post.Save(db)
	checkErr(err)
	r.JSON(http.StatusOK, post)
}

func deleteTweet(
	db gorp.SqlExecutor,
	params martini.Params,
	r render.Render) {
    post, err := GetPost(db, params["id"])
	checkErr(err)
	if post == nil {
		r.JSON(http.StatusNotFound, "NotFound")
		return
	}
    err = post.Delete(db)
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
