package services

import (
	"context"
	"database/sql"
	"gin-app/dto"
	"gin-app/errors"
	"gin-app/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
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

func (s *UserService) CreateUser(ctx context.Context, createUser dto.CreateUserDTO) (*models.User, error) {
	// check conflict
	cEmail, cName, err := s.repo.CheckConflict(ctx, createUser.Name, createUser.Email)
	// handle error if error not nil
	if err != nil {
		return nil, errors.NewInternalServerError("Something Went Wrong")
	}
	// if conflict found, return conflict error
	if cEmail || cName {
		return nil, errors.NewConflictError("Name or Email Already Exists")
	}
	// hashed password logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalServerError("failed to hash password")
	}

	// 4️⃣ Insert to DB
	user, err := s.repo.Create(ctx, &models.User{
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, errors.NewInternalServerError("failed to create user")
	}

	return user, nil
}
