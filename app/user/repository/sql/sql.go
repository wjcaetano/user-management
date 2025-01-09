package sql

import (
	entity "user-management/app/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	Writer *WriterSQL
	Reader *ReaderSQL
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Writer: NewWriterSQL(db),
		Reader: NewReaderSQL(db),
	}
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.Writer.Create(user)
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.Writer.Update(user)
}

func (r *UserRepository) Delete(id int) error {
	return r.Writer.Delete(id)
}

func (r *UserRepository) FindByID(id int) (*entity.User, error) {
	return r.Reader.FindByID(id)
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	return r.Reader.FindByEmail(email)
}

func (r *UserRepository) FindAll() ([]*entity.User, error) {
	return r.Reader.FindAll()
}

func (r *UserRepository) FindAllWithLimit(limit int) ([]*entity.User, error) {
	return r.Reader.FindAllWithLimit(limit)
}
