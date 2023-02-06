package controller

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllCars(w http.ResponseWriter, r *http.Request) {
	var cars []models.Car
	database.DB.Find(&cars)
	json.NewEncoder(w).Encode(cars)
}

func ReturnOneCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var car models.Car
	database.DB.First(&car, id)
	json.NewEncoder(w).Encode(car)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var newCar models.Car
	/*
	 *	Essa função abaixo é usada para decodificar os dados de um corpo de solicitação HTTP (r.Body) no formato JSON
	 *	e armazená-los em uma variável (newCar). O Pacote json fornece a função newDecoder para criar um novo decodificador
	 *	JSON e o método Decode para decodificar os dados JSON em um valor fornecido.
	 *
	 *	newCar precisa ser uma referência de uma estrutura ou tipo de dados com campos anotados com tags JSON para que
	 *	os dados possam ser mapeados para esses campos ao decodificar o JSON
	 */
	json.NewDecoder(r.Body).Decode(&newCar)

	/*
	 *	O método Create aceita uma referência de um tipo de dados e usa reflexão para inserir na tabela.
	 *
	 *	O processo de reflexão verifica as tags de campos da estrutura `newCar` para determinar como os dados devem ser
	 *	mapeados para as colunas da tabela do banco de dados. Depois, a linha é inserida na tabela usando INSERT.
	 */
	database.DB.Create(&newCar)
	// codificar para escrever no json
	json.NewEncoder(w).Encode(newCar)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var car models.Car
	database.DB.Delete(&car, id)
	json.NewEncoder(w).Encode(car)
}

func EditCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var car models.Car
	database.DB.First(&car, id)
	json.NewDecoder(r.Body).Decode(&car)
	database.DB.Save(&car)
	json.NewEncoder(w).Encode(car)
}
