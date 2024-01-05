package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user by id"))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
