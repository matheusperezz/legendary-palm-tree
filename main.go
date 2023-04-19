package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
	"os"
)

func main() {
	os.Setenv("DATABASE_PASSWORD", "Dofl$mingo432")
	database.ConnectToDatabase()
	fmt.Println("Starting Car Chase API...")
	routes.HandleRequest()
}
