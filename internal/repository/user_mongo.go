package repository

import (
	"context"
	"time"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *catMongo) GetUser(username, password string) (model.User, error) {
	var user model.User
	//var id string

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//filter := bson.M{"_id": id}
	users := bson.M{
		"username": username,
		"password": password,
	}
	err:= usersTable.FindOne(ctx, users).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, err
}
func (c *catMongo) CreateUser(user model.User) (int, error) {

	var id int

	_, err := usersTable.InsertOne(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil
}
