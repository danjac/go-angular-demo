package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func renderJSON(w http.ResponseWriter, status int, value interface{}) {
	w.WriteHeader(status)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(value)
}

func getTweets(w http.ResponseWriter, r *http.Request) {
	posts, err := GetPosts()
	checkErr(err)
	renderJSON(w, http.StatusOK, posts)
}

func addTweet(w http.ResponseWriter, r *http.Request) {

	post := &Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	checkErr(err)

	errors := post.Validate(r)

	if errors.Count() > 0 {
		renderJSON(w, http.StatusConflict, errors)
		return
	}

	err = post.Save()
	checkErr(err)
	renderJSON(w, http.StatusOK, post)
}

func deleteTweet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post, err := GetPost(vars["id"])
	checkErr(err)
	if post == nil {
		renderJSON(w, http.StatusNotFound, "NotFound")
		return
	}
	err = post.Delete()
	checkErr(err)
	renderJSON(w, http.StatusOK, "Deleted")
}

func SetupRoutes() *mux.Router {

	r := mux.NewRouter()

    // API
    s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/", getTweets).Methods("GET")
	s.HandleFunc("/", addTweet).Methods("POST")
	s.HandleFunc("/{id}", deleteTweet).Methods("DELETE")

	// STATIC FILES
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}
