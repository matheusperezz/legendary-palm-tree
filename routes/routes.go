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
	r.HandleFunc("/api/carros", controller.AllCars).Methods("Get")
	r.HandleFunc("/api/carros/{id}", controller.ReturnOneCar).Methods("Get")
	r.HandleFunc("/api/carros", controller.CreateCar).Methods("Post")
	r.HandleFunc("/api/carros/{id}", controller.DeleteCar).Methods("Delete")
	r.HandleFunc("/api/carros/{id}", controller.EditCar).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}
