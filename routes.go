package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func PostListHandler(ctx *RequestContext) {
	posts, err := GetPosts()
	if err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.RenderJSON(http.StatusOK, posts)
}

func CreatePostHandler(ctx *RequestContext) {

	post := &Post{}
	if err := ctx.DecodeJSON(post); err != nil {
		ctx.HandleError(err)
		return
	}

	if errors := post.Validate(); errors.Count() > 0 {
		ctx.RenderJSON(http.StatusConflict, errors)
		return
	}

	if err := post.Save(); err != nil {
		ctx.HandleError(err)
		return
	}
	ctx.RenderJSON(http.StatusOK, post)
}

func DeletePostHandler(ctx *RequestContext) {
	post, err := GetPost(ctx.Var("id"))
	if err != nil {
		ctx.HandleError(err)
		return
	}
	if post == nil {
		ctx.HandleNotFound("Post not found")
		return
	}
	if err := post.Delete(); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.HandleOK("Post deleted")
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
