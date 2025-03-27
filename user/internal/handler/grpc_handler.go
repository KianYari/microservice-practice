package handler

import (
	"context"

	pb "github.com/kianyari/microservice-practice/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
}

func NewGRPCHandler(
	grpcServer *grpc.Server,
) {
	handler := &grpcHandler{}
	pb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *grpcHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}

func (h *grpcHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Token: "sample token"}, nil
}
