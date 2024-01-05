package main

import (
	"fmt"
	"net/http"

	"github.com/LucioSchiavoni/go-postgres/db"
	"github.com/LucioSchiavoni/go-postgres/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Corriendo server de go")

	db.DBConnection()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":3000", r)
}
