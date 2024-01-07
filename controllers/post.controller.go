package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LucioSchiavoni/go-postgres/db"
	"github.com/LucioSchiavoni/go-postgres/models"
	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	createPost := db.DB.Create(&post)
	err := createPost.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	}

	json.NewEncoder(w).Encode(&post)

}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json/application")
	var post []models.Post
	db.DB.Find(&post)
	json.NewEncoder(w).Encode(&post)

}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json/application")
	params := mux.Vars(r)
	var post models.Post
	db.DB.First(&post, params["id"])

	if post.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Post not found"))
		return
	}

	json.NewEncoder(w).Encode(&post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	db.DB.First(&post, params["id"])
	if post.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Delete(&post)
	w.WriteHeader(http.StatusOK)

}
