package main

import (
	"fmt"
	"net/http"

	"github.com/LucioSchiavoni/go-postgres/db"
	"github.com/LucioSchiavoni/go-postgres/models"
	"github.com/LucioSchiavoni/go-postgres/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Corriendo server de go")

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Post{})
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/createUsers", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/deleteUser", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/posts", routes.GetPostsHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", routes.GetPostByIdHandler).Methods("GET")
	r.HandleFunc("/createPost", routes.CreatePostHandler).Methods("POST")
	r.HandleFunc("/deletePost", routes.DeletePostHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
