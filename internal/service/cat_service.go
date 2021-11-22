package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hyperstone1/TestCRUD/internal/model"
	"github.com/hyperstone1/TestCRUD/internal/repository"
	"github.com/pkg/errors"
)

var cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
var ctx = context.Background()

type CatService struct {
	rep *repository.Repositories
}

func init() {
	cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func New(rep *repository.Repositories) (*CatService, error) {
	if rep == nil {
		return nil, errors.New("service.New error")
	}
	return &CatService{rep}, nil
}

func (c *CatService) Create(cat model.Cat) error {
	err := c.rep.Cat.Create(cat)
	if err != nil {
		return err
	}
	return nil
}

func (c *CatService) Get() (model.Cats, error) {

	cats, err := c.rep.Cat.Get()
	if err != nil {
		return nil, errors.Wrap(err, "service.Get error")
	}
	return cats, nil
}

func (c *CatService) Delete(id string) error {

	_, err := c.rep.Cat.Get()
	if err != nil {
		return err
	}
	err = c.rep.Cat.Delete(id)
	if err != nil {
		return errors.Wrap(err, "service.Delete error")
	}

	return nil

}

func (c *CatService) Update(cat model.Cat) error {

	err := c.rep.Cat.Update(cat, cat.Id)

	if err != nil {
		return errors.Wrap(err, "service.Update error")
	}

	return nil
}

func (c *CatService) GetById(id string) (*model.Cat, error) {
	//var err error

	if t := cache.Get(ctx, id); t.Val() != "" {
		fmt.Println(t)
		b, err := t.Bytes()
		if err != nil {
			return nil, errors.Wrap(err, "service.GetById error")
		}
		cat := model.FromJson(b)
		return &cat, nil
	}
	cat, err := c.rep.Cat.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetById error")
	}

	cacheErr := cache.Set(ctx, id, model.ToJson(cat), 10*time.Minute)
	fmt.Println(cacheErr)
	return cat, err

}
