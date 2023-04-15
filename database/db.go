package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	connectionString := "user=postgres dbname=cars password=root host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		fmt.Println(err)
		return
	}
}
