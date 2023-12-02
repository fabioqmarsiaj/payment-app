package controllers

import (
	"api/models"
	"api/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro	= json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	userToCreate := models.User{ID: primitive.NewObjectID(),Name: user.Name, Email: user.Email, Password: user.Password, CreatedAt: time.Now()}

	userCreated, err := repositories.Create(userToCreate)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created with ID %s", userCreated.InsertedID)))

}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	allUsers, err := repositories.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	allUsersJson, _ := json.Marshal(allUsers)

	w.Write([]byte(fmt.Sprintf("All Users: %s", allUsersJson)))
}