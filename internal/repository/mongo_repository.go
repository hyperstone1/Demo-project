package repository

import (
	"context"
	"time"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type catMongo struct {
	client *mongo.Client
	cats   *mongo.Collection
	users  *mongo.Collection
}

const database = "cats"

var (
	ctx = context.Background()
)

func NewMongo(client *mongo.Client) (*catMongo, error) {
	err := client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	cats := client.Database(database).Collection("cats")
	users := client.Database(database).Collection("users")
	return &catMongo{
		client: client,
		cats:   cats,
		users:  users}, nil
}

func (c *catMongo) Create(cat model.Cat) error {

	_, err := c.cats.InsertOne(ctx, cat)

	if err != nil {
		return err
	}

	return nil

}

func (c *catMongo) Get() (model.Cats, error) {

	var cats model.Cats

	filter := bson.D{}

	cur, err := c.cats.Find(ctx, filter)

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

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name": cat.Name,
			"age":  cat.Age,
		},
	}

	_, err := c.cats.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (c *catMongo) Delete(id string) error {

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}

	_, err := c.cats.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

func (c *catMongo) GetById(id string) (*model.Cat, error) {
	var cat model.Cat
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}

	err := c.cats.FindOne(ctx, filter).Decode(&cat)

	if err != nil {
		return &cat, err
	}

	return &cat, nil
}
