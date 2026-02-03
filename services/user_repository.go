package services

import (
	"context"
	"gin-app/models"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
	CheckConflict(ctx context.Context, name, email string) (bool, bool, error)
	Create(
		ctx context.Context,
		user *models.User,
	) (*models.User, error)
}
