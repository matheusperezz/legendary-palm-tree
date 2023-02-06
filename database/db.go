package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// No caso de erros com docker, eu tentei:

/*
 *	1 - Trocar o nome do user,database e password
 *
 *	2 - reabrir o docker
 *
 *	3 - Reinstanciar o docker
 *
 *	4 - Recriar o banco de dados
 *
 *
 */

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	connectionString := "user=postgres dbname=cars password=Dofl$mingo432 host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		fmt.Println(err)
		return
	}
}
