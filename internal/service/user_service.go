package service

import (
	"github.com/progRayhan/CBS-360/internal/api/dto"
	"github.com/progRayhan/CBS-360/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
