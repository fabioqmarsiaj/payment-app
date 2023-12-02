package repositories

import (
	"api/db"
	"api/dto"
	"api/mappers"
	"api/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// insert an user into DB
func Create(user models.User) (*mongo.InsertOneResult, error) {
	
	client, err := db.Connect()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("payments").Collection("users")
	
	userCreated, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
	}
	return userCreated, nil
}

func GetAll() ([]dto.UserDTO, error) {
	client, err := db.Connect()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("payments").Collection("users")

	allUsers, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var allUsersList []models.User

	for allUsers.Next(context.TODO()){
		var result models.User
		// If err is not null, will populate err, otherwise will change de allUsers fields
		if err := allUsers.Decode(&result); err != nil {
			log.Fatal("Failed to decode", err)
		}
		allUsersList = append(allUsersList, result)
	}

	var allUsersListDTO []dto.UserDTO

	for _, user := range allUsersList {
		userDTO := mappers.UserToUserDTO(&user)
		allUsersListDTO = append(allUsersListDTO, *userDTO)
	}

	return allUsersListDTO, err
}