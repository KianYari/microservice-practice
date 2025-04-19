package main

import (
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/kianyari/microservice-practice/common/api"
	"github.com/kianyari/microservice-practice/gateway/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	userConn, err := grpc.NewClient("user_service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer userConn.Close()
	log.Println("Connected to user gRPC server on port :50051")

	taskConn, err := grpc.NewClient("task_service:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer taskConn.Close()
	log.Println("Connected to task gRPC server on port :50052")

	userClient := pb.NewUserServiceClient(userConn)
	taskClient := pb.NewTaskServiceClient(taskConn)
	jwtClient := pb.NewJWTServiceClient(userConn)

	ginEngine := gin.Default()

	userHandler := handler.NewUserHandler(userClient)
	taskHandler := handler.NewTaskHandler(taskClient, jwtClient)

	userHandler.RegisterRoutes(ginEngine)
	taskHandler.RegisterRoutes(ginEngine)

	ginEngine.Run(":8080")
}
