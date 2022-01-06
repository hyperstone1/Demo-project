package service

//С монго должны работать тесты
import (
	"testing"

	"github.com/hyperstone1/TestCRUD/internal/model"
	"github.com/hyperstone1/TestCRUD/internal/service"
)

type Serv struct {
	serv *service.CatService
}

func (s *Serv) TestCreateCat(t *testing.T) {
	cat := model.Cat{
		Name: "Jonh Milky",
		Age:  16,
	}

	err := s.serv.Create(cat)

	if err != nil {
		t.Error("Create not successed")
	} else {
		t.Log("Create successed")
	}
}

/*func (s *Serv) TestGet(t *testing.T, id uuid.UUID) {
	cats, err := s.serv.GetCats()

	if err != nil {
		t.Error("not successed")
		t.Fail()
	}

	if len(cats.Id) == 0 {
		t.Error("fadfadfad")
		t.Fail()
	} else {
		t.Log("fdgaga")
	}

}*/

func (s *Serv) TestUpdate(t *testing.T) {
	cat := model.Cat{
		Name: "Milky Way",
		Age:  44,
	}
	err := s.serv.Update(cat)

	if err != nil {
		t.Error("Create not successed")
	} else {
		t.Log("Create successed")
	}
}

/*func (s *Serv) TestDelete(t *testing.T) {
	err := s.serv.DeleteCat("44537769-9e9a-4ae4-99cf-523b6e351a34")

	if err != nil {
		t.Error("afagdag")
		t.Fail()
	} else {
		t.Log("dfadga")
	}
}*/
/*package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	host     = "localhost"
	port     = 27017
	database = "cats"
)

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
*/
