package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderJSON(w http.ResponseWriter, status int, value interface{}) {
	w.WriteHeader(status)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(value)
}

func PostListHandler(w http.ResponseWriter, r *http.Request) error {
	posts, err := GetPosts()
	if err != nil {
		return err
	}
	renderJSON(w, http.StatusOK, posts)
	return nil
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) error {

	post := &Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		return err
	}

	errors := post.Validate(r)

	if errors.Count() > 0 {
		renderJSON(w, http.StatusConflict, errors)
		return nil
	}

	err = post.Save()
	if err != nil {
		return err
	}
	renderJSON(w, http.StatusOK, post)
	return nil
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	post, err := GetPost(vars["id"])
	if err != nil {
		return err
	}
	if post == nil {
		renderJSON(w, http.StatusNotFound, "NotFound")
		return nil
	}
	err = post.Delete()
	if err != nil {
		return err
	}
	renderJSON(w, http.StatusOK, "Deleted")
	return nil
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
