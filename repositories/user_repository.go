package repositories

import (
	"ticketingapp/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindByID(id uint) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
	FindByIDRaw(id uint) (*entity.User, error)
	UpdateProfilePicture(id uint, path string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByIDRaw(id uint) (*entity.User, error) {
	return r.FindByID(id)
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) UpdateProfilePicture(id uint, path string) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Update("profile_picture", path).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
