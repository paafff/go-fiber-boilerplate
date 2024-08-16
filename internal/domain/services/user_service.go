package services

import (
    "fiber-app-paafff/internal/domain/models"
    "fiber-app-paafff/internal/repositories"
)

type UserService struct {
    userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
    return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
    return s.userRepository.GetUser(id)
}