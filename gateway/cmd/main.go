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
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to gRPC server on port :50051")

	userClient := pb.NewUserServiceClient(conn)

	gineEngine := gin.Default()

	userHandler := handler.NewUserHandler(userClient)

	userHandler.RegisterRoutes(gineEngine)

	gineEngine.Run(":8080")
}
