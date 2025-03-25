package service

import (
	"errors"

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
	_, exist := s.userRepository.GetUserByEmail(registerInfo.Email)
	if exist {
		return errors.New("user already exists")
	}
	user := &model.User{
		Email:    registerInfo.Email,
		Password: registerInfo.Password,
	}
	return s.userRepository.CreateUser(user)
}
func (s *userService) Login(loginInfo *dto.LoginRequest) (string, error) {
	user, exist := s.userRepository.GetUserByEmail(loginInfo.Email)
	if !exist {
		return "", errors.New("user not found")
	}
	if user.Password != loginInfo.Password {
		return "", errors.New("invalid password")
	}
	token, err := s.jwtService.GenerateToken(user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}
