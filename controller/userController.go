package controller

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	_ "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"net/http"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	CreatedUser(database.DB, newUser.Name, newUser.Email, newUser.Phone, newUser.Password)
	json.NewEncoder(w).Encode("User registred")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreatedUser(db *gorm.DB, name, email, phone, password string) (*models.User, error) {
	user := &models.User{Name: name, Email: email, Phone: phone}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
