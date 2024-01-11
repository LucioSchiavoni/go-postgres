package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LucioSchiavoni/go-postgres/db"
	"github.com/LucioSchiavoni/go-postgres/models"
	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error de formato FormData"})
		return
	}
	title := r.FormValue("title")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	userIDStr := r.FormValue("userId")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Error al parsear el precio", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Error al parsear el UserID", http.StatusBadRequest)
		return
	}

	newPost := models.Post{
		Title:       title,
		Description: description,
		Price:       price,
		UserID:      uint(userID),
	}
	file, _, err := r.FormFile("image_post")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al obtener la imagen"})
		return
	}
	defer file.Close()
	imagePath, err := UploadFile(w, r, "image_post")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al subir la imagen"})
		return
	}

	newPost.ImagePost = imagePath

	createPost := db.DB.Create(&newPost)
	err = createPost.Error
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

func GetPostByDescription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json/application")
	params := mux.Vars(r)
	userID := params["user_id"]
	description := params["description"]
	var post models.Post
	db.DB.First(&post, "user_id= ? AND description= ?", userID, description)
	if post.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Post not found"))
		return
	}

	json.NewEncoder(w).Encode(&post)
}

func GetPostByIdUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userID := params["user_id"]
	var posts []models.Post
	if err := db.DB.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&posts)
}

func GetPostByTitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]
	title := params["title"]
	var posts []models.Post
	if err := db.DB.Where("user_id = ? AND title = ?", userID, title).Find(&posts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&posts)
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
