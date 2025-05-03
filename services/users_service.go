package services

import (
	"ticketingapp/entity"
	"ticketingapp/repositories"
)

type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	UpdateUser(id uint, update entity.User) (*entity.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]entity.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, update entity.User) (*entity.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = update.Name
	user.Email = update.Email
	err = s.repo.Update(user)
	return user, err
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
