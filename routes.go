package main

import (
	"github.com/gorilla/mux"
	"net/http"
    "encoding/json"
)

func RenderJSON(w http.ResponseWriter, status int, value interface{}) {
	w.WriteHeader(status)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(value)
}

func Render(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func HandleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}


func PostListHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := GetPosts()
	if err != nil {
		HandleError(w,err)
		return
	}
	RenderJSON(w, http.StatusOK, posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	post := &Post{}
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		HandleError(w,err)
		return
	}

	if errors := post.Validate(); errors.Count() > 0 {
		RenderJSON(w, http.StatusConflict, errors)
		return
	}

	if err := post.Save(); err != nil {
		HandleError(w,err)
		return
	}
	RenderJSON(w, http.StatusCreated, post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	post, err := GetPost(mux.Vars(r)["id"])
	if err != nil {
		HandleError(w,err)
		return
	}
	if post == nil {
		Render(w, http.StatusNotFound, "Post not found")
		return
	}
	if err := post.Delete(); err != nil {
		HandleError(w,err)
		return
	}

	Render(w, http.StatusOK, "Post deleted")
}

func SetupRoutes() *mux.Router {

	r := mux.NewRouter()

	// API
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/", PostListHandler).Methods("GET")
	s.HandleFunc("/", CreatePostHandler).Methods("POST")
	s.HandleFunc("/{id}", DeletePostHandler).Methods("DELETE")

	// STATIC FILES
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}
