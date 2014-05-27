package main 

import (
	"encoding/json"
    "net/http"
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
