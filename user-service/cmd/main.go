package main

import (
	"fmt"
	"log"

	handler "github.com/kianyari/microservice-practice/user-service/internal/handler"
	model "github.com/kianyari/microservice-practice/user-service/internal/model"
	repository "github.com/kianyari/microservice-practice/user-service/internal/repository"
	service "github.com/kianyari/microservice-practice/user-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/kianyari/microservice-practice/user-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(
		&model.User{},
	)

	userRepository := repository.NewUserRepository(db)
	jwtService := service.NewJWTService(cfg.JWTSecret)
	userService := service.NewUserService(userRepository, jwtService)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
