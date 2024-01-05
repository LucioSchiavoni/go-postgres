package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el .env")
	}

}

func DBConnection() *gorm.DB {
	var DNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Sprintlf("Valor %s", DNS)
	DB, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")

	}
	return DB
}
