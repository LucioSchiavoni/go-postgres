package routes

import (
	"github.com/LucioSchiavoni/go-postgres/controllers"
	"github.com/gorilla/mux"
)

func PostRoutes(router *mux.Router) {
	router.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/post/{id}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/post", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/description/{user_id}/{description}", controllers.GetPostByDescription).Methods("GET")
	router.HandleFunc("/usuario/{user_id}/posts", controllers.GetPostByIdUser).Methods("GET")
	// el usuario con user_id = ? , Find &posts (buscar todos sus posts con su id de usuario)
}
