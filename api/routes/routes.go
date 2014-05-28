package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/danjac/angular-react-compare/api/models"
	"github.com/danjac/angular-react-compare/api/render"
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

func NewRouter(secretKey string) http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/", getPosts).Methods("GET")
	r.HandleFunc("/", createPost).Methods("POST")
	r.HandleFunc("/{id}", deletePost).Methods("DELETE")

	return NewCSRF(secretKey, r)
}
