package routes

import (
	"encoding/json"
	"github.com/danjac/go-angular-demo/api/models"
	"github.com/danjac/go-angular-demo/api/render"
	"github.com/gorilla/mux"
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

	if result := post.Validate(); !result.OK {
		render.JSON(w, http.StatusConflict, result)
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

func Init(r *mux.Router) {
	r.HandleFunc("/", getPosts).Methods("GET")
	r.HandleFunc("/", createPost).Methods("POST")
	r.HandleFunc("/{id}", deletePost).Methods("DELETE")
}
