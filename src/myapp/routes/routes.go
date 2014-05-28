package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"myapp/models"
	"myapp/render"
	"net/http"
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetPosts()
	if err != nil {
		render.Error(w, err)
		return
	}
	render.JSON(w, http.StatusOK, posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {

	post := &models.Post{}
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		render.Error(w, err)
		return
	}

	if errors := post.Validate(); errors.Count() > 0 {
		render.JSON(w, http.StatusConflict, errors)
		return
	}

	if err := post.Save(); err != nil {
		render.Error(w, err)
		return
	}
	render.JSON(w, http.StatusCreated, post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	post, err := models.GetPost(mux.Vars(r)["id"])
	if err != nil {
		render.Error(w, err)
		return
	}
	if post == nil {
		render.Status(w, http.StatusNotFound, "No post found")
		return
	}
	if err := post.Delete(); err != nil {
		render.Error(w, err)
		return
	}

	render.Status(w, http.StatusOK, "Post deleted")
}

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	// API
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/", getPosts).Methods("GET")
	s.HandleFunc("/", createPost).Methods("POST")
	s.HandleFunc("/{id}", deletePost).Methods("DELETE")

	// STATIC FILES
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	return r
}
