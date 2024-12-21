package sql

import (
	entity "user-management/app/user"

	"gorm.io/gorm"
)

type WriterSQL struct {
	DB *gorm.DB
}

func NewWriterSQL(db *gorm.DB) *WriterSQL {
	return &WriterSQL{DB: db}
}

func (r *WriterSQL) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *WriterSQL) Update(user *entity.User) error {
	return r.DB.Save(user).Error
}

func (r *WriterSQL) Delete(id int) error {
	return r.DB.Delete(&entity.User{}, id).Error
}
