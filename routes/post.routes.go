package routes

import "net/http"

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all post"))
}

func GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get post by id"))
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create post"))
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete post"))
}
