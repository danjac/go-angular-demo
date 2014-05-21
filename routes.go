package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func PostListHandler(ctx *RequestContext) {
	posts, err := GetPosts()
	if err != nil {
		ctx.RenderError(err)
		return
	}
	ctx.RenderJSON(http.StatusOK, posts)
}

func CreatePostHandler(ctx *RequestContext) {

	post := &Post{}
	if err := ctx.DecodeJSON(post); err != nil {
		ctx.RenderError(err)
		return
	}

	errors := post.Validate(ctx.Request)

	if errors.Count() > 0 {
		ctx.RenderJSON(http.StatusConflict, errors)
		return
	}

	if err := post.Save(); err != nil {
		ctx.RenderError(err)
		return
	}
	ctx.RenderJSON(http.StatusOK, post)
}

func DeletePostHandler(ctx *RequestContext) {
	post, err := GetPost(ctx.Var("id"))
	if err != nil {
		ctx.RenderError(err)
		return
	}
	if post == nil {
		ctx.RenderJSON(http.StatusNotFound, "NotFound")
		return
	}
	if err := post.Delete(); err != nil {
		ctx.RenderError(err)
		return
	}

	ctx.RenderJSON(http.StatusOK, "Deleted")
}

func SetupRoutes() *mux.Router {

	r := mux.NewRouter()

	// API
	s := r.PathPrefix("/api").Subrouter()

	s.Handle("/", AppHandler(PostListHandler)).Methods("GET")
	s.Handle("/", AppHandler(CreatePostHandler)).Methods("POST")
	s.Handle("/{id}", AppHandler(DeletePostHandler)).Methods("DELETE")

	// STATIC FILES
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}
