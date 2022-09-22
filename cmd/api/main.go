package main

import (
	"os"

	"github.com/koba1108/go-mongodb/internals/client"
	"github.com/koba1108/go-mongodb/internals/handler"
	"github.com/koba1108/go-mongodb/internals/infrastructure/database"
	"github.com/koba1108/go-mongodb/internals/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	mongodb, err := client.NewMongoDBClient()
	if err != nil {
		panic(err)
	}

	// wire化できそう
	articleRepo := database.NewArticleRepository(mongodb)
	articleUsecase := usecase.NewArticleUsecase(articleRepo)
	articleHandler := handler.NewArticleHandler(articleUsecase)

	userRepo := database.NewUserRepository(mongodb)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	playerRepo := database.NewPlayerRepository(mongodb)
	playerUsecase := usecase.NewPlayerUsecase(playerRepo)
	playerHandler := handler.NewPlayerHandler(playerUsecase)

	videoRepo := database.NewVideoRepository(mongodb)
	videoUsecase := usecase.NewVideoUsecase(videoRepo, playerRepo)
	videoHandler := handler.NewVideoHandler(videoUsecase)

	sampleDataHandler := handler.NewSampleDataHandler(videoRepo, playerRepo)

	e := echo.New()
	e.Debug = true

	sampleData := e.Group("sample-data")
	{
		sampleData.POST("", sampleDataHandler.Create)
	}

	apiV1 := e.Group("/api/v1")
	{
		apiV1Articles := apiV1.Group("/articles")
		{
			apiV1Articles.GET("", articleHandler.Find)
		}
		apiV1User := apiV1.Group("/users")
		{
			apiV1User.GET("", userHandler.List)
			apiV1User.GET("/:id", userHandler.GetByID)
		}
		apiV1Player := apiV1.Group("/players")
		{
			apiV1Player.GET("", playerHandler.List)
			apiV1Player.GET("/:id", playerHandler.GetByID)
			apiV1Player.POST("", playerHandler.Create)
			apiV1Player.PUT("/:id", playerHandler.Update)
			apiV1Player.DELETE("/:id", playerHandler.Delete)
		}
		apiV1Video := apiV1.Group("/videos")
		{
			apiV1Video.GET("", videoHandler.List)
			apiV1Video.GET("/:id", videoHandler.GetByID)
			apiV1Video.POST("", videoHandler.Create)
			apiV1Video.PUT("/:id", videoHandler.Update)
			apiV1Video.DELETE("/:id", videoHandler.Delete)
		}
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
