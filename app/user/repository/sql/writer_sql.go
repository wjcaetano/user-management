package sql

import (
	"gorm.io/gorm"
	entity "user-management/app/user"
)

type WriterSql struct {
	DB *gorm.DB
}

func NewWriterSQL(db *gorm.DB) *WriterSql {
	return &WriterSql{DB: db}
}

func (r *WriterSql) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *WriterSql) Update(user *entity.User) error {
	return r.DB.Save(user).Error
}

func (r *WriterSql) Delete(id int) error {
	return r.DB.Delete(&entity.User{}, id).Error
}
