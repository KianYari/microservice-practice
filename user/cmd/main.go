package main

import (
	"fmt"
	"log"
	"net"

	handler "github.com/kianyari/microservice-practice/user-service/internal/handler"
	model "github.com/kianyari/microservice-practice/user-service/internal/model"
	repository "github.com/kianyari/microservice-practice/user-service/internal/repository"
	service "github.com/kianyari/microservice-practice/user-service/internal/service"
	"google.golang.org/grpc"

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
	userService.Register("kian", "password")

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer()
	handler.NewGRPCHandler(grpcServer)

	log.Println("gRPC server is running on port :50051")

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
