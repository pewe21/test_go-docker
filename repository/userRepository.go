package repository

import (
	"github.com/pewe21/go-docker/schema"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id int) (*schema.User, error)
	Create(user *schema.User) error
	Update(id int, user *schema.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByID(id int) (*schema.User, error) {
	var user schema.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *schema.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(id int, user *schema.User) error {
	if err := r.db.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&schema.User{}).Error; err != nil {
		return err
	}
	return nil
}
