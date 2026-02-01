package services

import (
	"context"
	"database/sql"
	"gin-app/errors"
	"gin-app/models"
	"gin-app/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	user, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, errors.NewInternalServerError("Something Went Wrong")
	}
	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("User Not Found")
		}
		return nil, errors.NewInternalServerError("Something Went Wrong")
	}
	return user, nil
}
