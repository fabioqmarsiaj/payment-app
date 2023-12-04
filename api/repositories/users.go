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

var coll *mongo.Collection

func init(){
	client, err := db.Connect()

	if err != nil{
		client.Disconnect(context.TODO())
		log.Fatal(err)
	}

	coll = client.Database("payments").Collection("users")
}

// insert an user into DB
func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	
	userCreated, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
	}
	return userCreated, nil
}

func GetAllUsers() ([]dto.UserDTO, error) {

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

/* func GetUserByName(name string) (models.User, error) {

	return nil, nil
} */