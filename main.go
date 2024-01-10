package main

import (
	"fmt"
	"net/http"

	"github.com/LucioSchiavoni/go-postgres/db"
	"github.com/LucioSchiavoni/go-postgres/models"
	"github.com/LucioSchiavoni/go-postgres/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Corriendo server de go")

	r := mux.NewRouter()
	db.DBConnection()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	if isDevelopment() {
		db.DB.AutoMigrate(models.User{})
		db.DB.AutoMigrate(models.Post{})
	}

	http.Handle("/", handlers.CORS(headersOk, originsOk, methodsOk)(r))

	routes.UserRoutes(r)
	routes.PostRoutes(r)

	http.ListenAndServe(":3000", nil)
}

func isDevelopment() bool {
	return false
}
