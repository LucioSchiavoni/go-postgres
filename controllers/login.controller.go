package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LucioSchiavoni/go-postgres/controllers"
	"github.com/LucioSchiavoni/go-postgres/db"
	"github.com/LucioSchiavoni/go-postgres/middlewares"
	"github.com/LucioSchiavoni/go-postgres/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginCredentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request body")
		return
	}

	fmt.Printf("The user request value %v", loginCredentials)

	var user models.User
	result := db.DB.Where("email = ?", loginCredentials.Email).First(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	if !controllers.CheckPasswordHash(loginCredentials.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	tokenString, err := middlewares.CreateToken(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error generando el token")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Error de Authorization")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := middlewares.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	fmt.Fprint(w, "Autenticado!")

}
