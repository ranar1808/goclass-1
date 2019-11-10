package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	ID     string `json:"id"`
	Tittle string `json:"title"`
	Body   string `json:"body"`
}

var posts []Post

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPosts).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePosts).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePosts).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
