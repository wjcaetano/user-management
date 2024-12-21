package sql

import (
	entity "user-management/app/user"

	"gorm.io/gorm"
)

type ReaderSQL struct {
	DB *gorm.DB
}

func NewReaderSQL(db *gorm.DB) *ReaderSQL {
	return &ReaderSQL{DB: db}
}

func (r *ReaderSQL) FindByID(id int) (*entity.User, error) {
	var user entity.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ReaderSQL) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("email = ?").First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ReaderSQL) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *ReaderSQL) FindAllWithLimit(limit int) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.DB.Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
