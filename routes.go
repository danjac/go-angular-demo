package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestContext struct {
    Response http.ResponseWriter
    Request *http.Request
    Vars map[string]string
}

func (ctx *RequestContext) Var (name string) string {
    return ctx.Vars[name]
}

func (ctx *RequestContext) RenderJSON(status int, value interface{}) {
	ctx.Response.WriteHeader(status)
	ctx.Response.Header().Add("content-type", "application/json")
	json.NewEncoder(ctx.Response).Encode(value)
}

func (ctx *RequestContext) RenderError(err error) {
	http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
}

func (ctx *RequestContext) DecodeJSON(value interface{}) error {
	return json.NewDecoder(ctx.Request.Body).Decode(value)
}

type appHandler func(ctx *RequestContext) 

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ctx := &RequestContext{Response: w, Request: r, Vars: mux.Vars(r)}
    fn(ctx)
}

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
	err := ctx.DecodeJSON(post)
	if err != nil {
		ctx.RenderError(err)
        return
	}

	errors := post.Validate(ctx.Request)

	if errors.Count() > 0 {
		ctx.RenderJSON(http.StatusConflict, errors)
        return
	}

	err = post.Save()
	if err != nil {
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
	}
	err = post.Delete()
	if err != nil {
        ctx.RenderError(err)
        return
	}

	ctx.RenderJSON(http.StatusOK, "Deleted")
}

func SetupRoutes() *mux.Router {

	r := mux.NewRouter()

	// API
	s := r.PathPrefix("/api").Subrouter()

	s.Handle("/", appHandler(PostListHandler)).Methods("GET")
	s.Handle("/", appHandler(CreatePostHandler)).Methods("POST")
	s.Handle("/{id}", appHandler(DeletePostHandler)).Methods("DELETE")

	// STATIC FILES
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}
