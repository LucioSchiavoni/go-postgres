package routes

import (
	"github.com/LucioSchiavoni/go-postgres/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	router.HandleFunc("/user", controllers.DeleteUserHandler).Methods("DELETE")
}
