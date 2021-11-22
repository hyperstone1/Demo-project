package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type catMongo struct {
	client *mongo.Client
}

var (
	host       = "localhost"
	port       = 27017
	database   = "cats"
	ctx        = context.Background()
	collection = GetCollection("cats")
	usersTable = GetCollection("users")
)

func New_Mongo(client *mongo.Client) (*catMongo, error) {

	return &catMongo{client}, nil
}

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%d", host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}
	return client.Database(database).Collection(collection)

}

func (c *catMongo) Create(cat model.Cat) error {

	var err error

	_, err = collection.InsertOne(ctx, cat)

	if err != nil {
		return err
	}

	return nil

}

func (c *catMongo) Get() (model.Cats, error) {

	var cats model.Cats

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var cat model.Cat
		err = cur.Decode(&cat)

		if err != nil {
			return nil, err
		}

		cats = append(cats, &cat)
	}
	return cats, nil
}

func (c *catMongo) Update(cat model.Cat, id string) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name": cat.Name,
			"age":  cat.Age,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (c *catMongo) Delete(id string) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}


func (c *catMongo) GetById(id string) (*model.Cat,error) {
	var cat model.Cat
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx,filter).Decode(&cat)
	
	if err != nil { 
		return &cat,err 
	}

	return &cat,nil
}
/*func (c *catMongo) GetUser(username, password string) (model.User, error) {
	var user model.User
	var id string

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{"_id": id}
	
	cur, err := usersTable.Find(ctx, filter)
	if err != nil {
		return user, err
	}
	for cur.Next(ctx) {
		var user model.User
		err = cur.Decode(&user)

		return user, err
	}
	return user, err
}

func (c *catMongo) GetUs(username, password string) (model.User,error) {
	var user model.User
	var id string
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(ctx, filter)
	cur, err := usersTable.Find(ctx, filter)
	for cur.Next(ctx) {
		var user model.User
		err = cur.Decode(&user)

		return user, err
	}
	return user, err
}*/
