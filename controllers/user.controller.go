package controllers

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"

	"github.com/LucioSchiavoni/go-postgres/db"
	"github.com/LucioSchiavoni/go-postgres/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UploadFile(w http.ResponseWriter, r *http.Request, fieldName string) (string, error) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile(fieldName)
	if err != nil {
		w.Write([]byte("Error al subir la imagen"))
		return "", err

	}

	defer file.Close()

	tempFile, err := os.CreateTemp("temp-images", "upload-*.png")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al crear el archivo temporal"})
		return "", err
	}

	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)

	if err != nil {
		fmt.Println(err)
	}

	return tempFile.Name(), nil

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error de formato FormData"})
		return
	}

	user.Username = r.FormValue("username")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	user.Address = r.FormValue("address")

	secret := os.Getenv("HASH_PWD")
	password := r.FormValue("password")

	hash, err := HashPassword(secret + password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al generar el hash de la contraseña"})
		return
	}

	user.Password = string(hash)

	file, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al obtener la imagen"})
		return
	}
	defer file.Close()

	imagePath, err := UploadFile(w, r, "image")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al subir la imagen"})
		return
	}

	user.Image = imagePath

	createUser := db.DB.Create(&user)
	err = createUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	user.Password = ""

	json.NewEncoder(w).Encode(&user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Posts").Find(&user.Posts)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
}
