package service

import (
	"errors"

	"github.com/kianyari/microservice-practice/user-service/internal/model"
	"github.com/kianyari/microservice-practice/user-service/internal/repository"
)

type UserService interface {
	Register(email string, password string) error
	Login(email string, password string) (string, error)
	GetUserByID(id uint) (*model.User, error)
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

func (s *userService) Register(email string, password string) error {
	_, exist := s.userRepository.GetUserByEmail(email)
	if exist {
		return errors.New("user already exists")
	}
	user := &model.User{
		Email:    email,
		Password: password,
	}
	return s.userRepository.CreateUser(user)
}
func (s *userService) Login(email string, password string) (string, error) {
	user, exist := s.userRepository.GetUserByEmail(email)
	if !exist {
		return "", errors.New("user not found")
	}
	if user.Password != password {
		return "", errors.New("invalid password")
	}
	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	user, exist := s.userRepository.GetUserByID(id)
	if !exist {
		return nil, errors.New("user not found")
	}
	return user, nil
}
