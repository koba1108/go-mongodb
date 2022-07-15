package main

import (
	"os"

	"github.com/koba1108/go-mongodb/internals/client"
	"github.com/koba1108/go-mongodb/internals/handler"
	"github.com/koba1108/go-mongodb/internals/infrastructure/database"
	"github.com/koba1108/go-mongodb/internals/usecase"
	"github.com/labstack/echo"
)

func main() {
	mongodb, err := client.NewMongoDBClient(os.Getenv("MONGODB_DATABASE"), os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}
	userRepo := database.NewUserRepository(mongodb)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	apiV1 := e.Group("/api/v1")
	apiV1User := apiV1.Group("users")
	{
		apiV1User.GET("/", userHandler.List)
		apiV1User.GET("/:userId", userHandler.GetByID)
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
