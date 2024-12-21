package repository

import (
	"user-management/app/user"
)

type UserRepository interface {
	Create(user *user.User) error
	Update(user *user.User) error
	Delete(id int) error
	FindByID(id int) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
	FindAll() ([]*user.User, error)
	FindAllWithLimit(limit int) ([]*user.User, error)
}
