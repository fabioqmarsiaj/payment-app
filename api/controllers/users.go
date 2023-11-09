package controllers

import (
	"api/models"
	"api/repositories"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// requestBody, erro := io.ReadAll(r.Body)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }

	// var user models.User
	// if erro	= json.Unmarshal(requestBody, &user); erro != nil {
	// 	log.Fatal(erro)
	// }

	userToCreate := models.User{Name: "Test", Email: "Test", Password: "Test", CreatedAt: time.Now()}

	userCreated, err := repositories.Create(userToCreate)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created with ID %s", userCreated.InsertedID)))

}