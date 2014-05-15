package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

func getTweets(r render.Render) {
	posts, err := GetPosts()
	checkErr(err)
	r.JSON(http.StatusOK, posts)
}

func addTweet(
	post Post,
	errors binding.Errors,
	r render.Render) {

	if errors.Count() > 0 {
		r.JSON(http.StatusConflict, errors)
		return
	}

	err := post.Save()
	checkErr(err)
	r.JSON(http.StatusOK, post)
}

func deleteTweet(
	params martini.Params,
	r render.Render) {
	post, err := GetPost(params["id"])
	checkErr(err)
	if post == nil {
		r.JSON(http.StatusNotFound, "NotFound")
		return
	}
	err = post.Delete()
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
