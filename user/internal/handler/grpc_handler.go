package handler

import (
	"context"

	pb "github.com/kianyari/microservice-practice/common/api"
	"github.com/kianyari/microservice-practice/user-service/internal/service"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewGRPCHandler(
	grpcServer *grpc.Server,
	userService service.UserService,
) {
	handler := &grpcHandler{
		userService: userService,
	}
	pb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *grpcHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := h.userService.Register(req.Email, req.Password)
	if err != nil {
		return &pb.RegisterResponse{Message: err.Error()}, err
	}
	return &pb.RegisterResponse{Message: "user registered successfully"}, nil
}

func (h *grpcHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.userService.Login(req.Email, req.Password)
	return &pb.LoginResponse{Token: token}, err
}
