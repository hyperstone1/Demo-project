package repository

import (
	"context"
	"os"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CatRepository interface {
	Get() (model.Cats, error)
	Create(cat model.Cat) error
	Update(cat model.Cat, id string) error
	Delete(id string) error
	GetById(id string) (*model.Cat, error)
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}
type Repositories struct {
	Cat CatRepository
}

func New(ctx context.Context) (*Repositories, error) {
	var rep Repositories
	var err error

	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found")
	}

	PgEnv, _ := os.LookupEnv("DB")

	if PgEnv == "PG" {
		conn, _ := pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/cats")
		rep.Cat, err = New_PG(conn)
		return &rep, err

	} else {
		client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/cats"))
		rep.Cat, err = New_Mongo(client)
		return &rep, err
	}

}
