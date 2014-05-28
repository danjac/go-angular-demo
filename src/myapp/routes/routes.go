package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"myapp/models"
	"myapp/render"
	"net/http"
)

func PostListHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetPosts()
	if err != nil {
		render.Error(w, err)
		return
	}
	render.JSON(w, http.StatusOK, posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

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

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	post, err := models.GetPost(mux.Vars(r)["id"])
	if err != nil {
		render.Error(w, err)
		return
	}
	if post == nil {
		render.Status(w, http.StatusNotFound)
		return
	}
	if err := post.Delete(); err != nil {
		render.Error(w, err)
		return
	}

	render.Status(w, http.StatusOK)
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
