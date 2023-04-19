package database

import (
	"api-go-rest/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	password := os.Getenv("DATABASE_PASSWORD")
	connectionString := "user=postgres dbname=carchase password=" + password + " host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		fmt.Println(err)
		return
	}
	err := DB.AutoMigrate(&models.User{}, &models.Car{})
	if err != nil {
		return
	}
}
