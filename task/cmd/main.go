package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/kianyari/microservice-practice/common/api"
	"github.com/kianyari/microservice-practice/task-service/config"
	handler "github.com/kianyari/microservice-practice/task-service/internal/handler"
	model "github.com/kianyari/microservice-practice/task-service/internal/model"
	repository "github.com/kianyari/microservice-practice/task-service/internal/repository"
	service "github.com/kianyari/microservice-practice/task-service/internal/service"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	l, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

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
		&model.Task{},
	)

	userConn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer userConn.Close()

	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort,
	)
	rabbitMQConn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()

	userClient := pb.NewUserServiceClient(userConn)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(userClient, taskRepository, rabbitMQConn)

	taskService.StartDeadlineChecker()

	grpcServer := grpc.NewServer()
	var serviceInterface service.TaskServiceInterface = taskService
	handler.NewGRPCHandler(grpcServer, serviceInterface)

	log.Println("gRPC server is running on port :50052")

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
