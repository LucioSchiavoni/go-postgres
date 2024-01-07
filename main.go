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

	if isDevelopment() {
		db.DB.AutoMigrate(models.User{})
		db.DB.AutoMigrate(models.Post{})
	}

	r := mux.NewRouter()

	routes.UserRoutes(r)
	routes.PostRoutes(r)

	http.ListenAndServe(":3000", r)
}

func isDevelopment() bool {
	return false
}
