package repository

import (
	"context"
	"database/sql"
	"gin-app/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.DB.QueryContext(ctx, `
	SELECT id, name, email 
	FROM users
	ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []models.User{}

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	row := r.DB.QueryRowContext(ctx, `
		SELECT id, name, email, created_at
		FROM users
		WHERE id = $1`, id)

	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) CheckConflict(ctx context.Context, name, email string) (bool, bool, error) {
	query := `
		SELECT
			COUNT(*) FILTER (WHERE email = $1) AS email_count,
			COUNT(*) FILTER (WHERE name = $2)  AS name_count
		FROM users
	`
	var emailCount, nameCount int

	err := r.DB.QueryRowContext(ctx, query, email, name).
		Scan(&emailCount, &nameCount)

	if err != nil {
		return false, false, err
	}

	return emailCount > 0, nameCount > 0, nil

}

func (r *UserRepository) Create(
	ctx context.Context,
	user *models.User,
) (*models.User, error) {

	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	err := r.DB.QueryRowContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
