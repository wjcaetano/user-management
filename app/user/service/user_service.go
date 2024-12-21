package service

import (
	"errors"
	entity "user-management/app/user"
	"user-management/app/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *entity.User) (*entity.User, error)
}

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) CreateUser(user *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user = &entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
