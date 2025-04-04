package repository

import (
	model "github.com/kianyari/microservice-practice/user-service/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, bool)
	GetUserByEmail(email string) (*model.User, bool)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*model.User, bool) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, false
	}
	return &user, true
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, bool) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, false
	}
	return &user, true
}
