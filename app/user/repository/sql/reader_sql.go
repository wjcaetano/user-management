package sql

import (
	"gorm.io/gorm"
	entity "user-management/app/user"
)

type ReaderSql struct {
	DB *gorm.DB
}

func NewReaderSql(db *gorm.DB) *ReaderSql {
	return &ReaderSql{DB: db}
}

func (r *ReaderSql) FindByID(id int) (*entity.User, error) {
	var user entity.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ReaderSql) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("email = ?").First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ReaderSql) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *ReaderSql) FindAllWithLimit(limit int) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.DB.Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
