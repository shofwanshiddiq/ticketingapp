package services

import (
	"ticketingapp/entity"
	"ticketingapp/repositories"
	"ticketingapp/utils"
)

type AuthService interface {
	Register(user *entity.User) (string, error)
	Login(email, password string) (string, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(user *entity.User) (string, error) {
	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	// Save user to DB
	if err := s.repo.Create(user); err != nil {
		return "", err
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// Check password using utility
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", err
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
