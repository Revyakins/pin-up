package app

import (
	"context"
	"go-app/pin-up/controllers"
	"go-app/pin-up/repositories/user_repo"
	"go-app/pin-up/services/user_service"
	"log"
	"github.com/joho/godotenv"
	"go-app/pin-up/configs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var r = gin.Default()

func Run()  {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.GetConfig()
	clientOptions := options.Client().ApplyURI(config.MongoDb.URI)
	mongoDB, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}

	userRepo := user_repo.NewUserRepository(mongoDB)
	userService := user_service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	/*
		===== Book Routes =====
	*/
	r.POST("/user/create", userController.CreateUser)
	r.Run()
}
