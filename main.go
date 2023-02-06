package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	database.ConnectToDatabase()
	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
