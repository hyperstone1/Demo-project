package main

import (
	"context"

	//"github.com/hyperstone1/TestCRUD/internal/cache"
	_ "github.com/hyperstone1/TestCRUD/docs"
	"github.com/hyperstone1/TestCRUD/internal/handler"
	"github.com/hyperstone1/TestCRUD/internal/repository"
	"github.com/hyperstone1/TestCRUD/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title CRUD API
// @version 1.0
// @description API Server for CRUD

// @host localhost:1328
// @BasePath /

// @secutiryDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	err := run()
	if err != nil {
		logrus.Fatal(err)
	}

}

func run() error {

	ctx := context.Background()
	repository, err := repository.New(ctx)
	if err != nil {
		return errors.Wrap(err, "Repository.New failed")
	}
	CatService, err := service.New(repository)
	if err != nil {
		return errors.Wrap(err, "Service.New error")
	}

	CatHandler, err := handler.New(CatService)
	if err != nil {
		return errors.Wrap(err, "Handler.New error")
	}
	e := echo.New()

	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", handler.UploadFile)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	auth := e.Group("/auth")
	{
		auth.POST("/signup/", CatHandler.SignUp)
		auth.POST("/signin/", CatHandler.SignIn)
	}

	api := e.Group("/api")
	{
		userGroup := api.Group("/cats", CatHandler.UserIdentity)
		{
			userGroup.GET("/", CatHandler.Get)
			userGroup.GET("/:id", CatHandler.GetById)
			userGroup.POST("/", CatHandler.Create)
			userGroup.PUT("/", CatHandler.Update)
			userGroup.DELETE("/:id", CatHandler.Delete)
		}
	}
	e.Logger.Fatal(e.Start(":1328"))

	return nil
}
