package routes

import (
	"api-go-rest/controller"
	"api-go-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	// Configuration
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)

	// Get Methods
	r.HandleFunc("/api/cars", controller.AllCars).Methods("Get")
	r.HandleFunc("/api/cars/{id}", controller.ReturnOneCar).Methods("Get")
	r.HandleFunc("/api/users", controller.AllUsers).Methods("Get")

	// Post Methods
	r.HandleFunc("/api/cars", controller.CreateCar).Methods("Post")
	r.HandleFunc("/api/users", controller.CreateNewUser).Methods("Post")

	// Delete Methods
	r.HandleFunc("/api/cars/{id}", controller.DeleteCar).Methods("Delete")

	// Put Methods
	r.HandleFunc("/api/cars/{id}", controller.EditCar).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
