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
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/cars", controller.AllCars).Methods("Get")
	r.HandleFunc("/api/cars/{id}", controller.ReturnOneCar).Methods("Get")
	r.HandleFunc("/api/cars", controller.CreateCar).Methods("Post")
	r.HandleFunc("/api/cars/{id}", controller.DeleteCar).Methods("Delete")
	r.HandleFunc("/api/cars/{id}", controller.EditCar).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
