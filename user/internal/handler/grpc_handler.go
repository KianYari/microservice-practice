package handler

import (
	"context"

	pb "github.com/kianyari/microservice-practice/common/api"
	"github.com/kianyari/microservice-practice/user-service/internal/service"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedJWTServiceServer
	userService service.UserService
	jwtService  service.JWTService
}

func NewGRPCHandler(
	grpcServer *grpc.Server,
	userService service.UserService,
	jwtService service.JWTService,
) {
	handler := &grpcHandler{
		userService: userService,
		jwtService:  jwtService,
	}
	pb.RegisterUserServiceServer(grpcServer, handler)
	pb.RegisterJWTServiceServer(grpcServer, handler)

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

func (h *grpcHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := h.userService.GetUserByID(uint(req.Id))
	if err != nil {
		return &pb.GetUserByIdResponse{Id: int32(user.ID), Email: user.Email}, err
	}
	return &pb.GetUserByIdResponse{
		Id:    int32(user.ID),
		Email: user.Email,
	}, nil
}

func (h *grpcHandler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	id, err := h.jwtService.ValidateToken(req.Token)
	return &pb.ValidateTokenResponse{Id: int32(id)}, err
}
