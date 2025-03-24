package service

import (
	"github.com/kianyari/microservice-practice/user-service/internal/dto"
	"github.com/kianyari/microservice-practice/user-service/internal/model"
	"github.com/kianyari/microservice-practice/user-service/internal/repository"
)

type UserService interface {
	Register(registerInfo *dto.RegisterRequest) error
	Login(loginInfo *dto.LoginRequest) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
	jwtService     JWTService
}

func NewUserService(
	userRepository repository.UserRepository,
	jwtService JWTService,
) UserService {
	return &userService{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (s *userService) Register(registerInfo *dto.RegisterRequest) error {
	_, err := s.userRepository.GetUserByEmail(registerInfo.Email)
	if err == nil {
		return err
	}
	user := &model.User{
		Email:    registerInfo.Email,
		Password: registerInfo.Password,
	}
	return s.userRepository.CreateUser(user)
}
func (s *userService) Login(loginInfo *dto.LoginRequest) (string, error) {
	user, err := s.userRepository.GetUserByEmail(loginInfo.Email)
	if err != nil {
		return "", err
	}
	if user.Password != loginInfo.Password {
		return "", err
	}
	token, err := s.jwtService.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
