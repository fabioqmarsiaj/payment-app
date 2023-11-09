package repositories

import (
	"api/db"
	"api/models"
	"context"
	"log"

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