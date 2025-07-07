package store

import (
	"context"
	"database/sql"

	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/models"
)

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{DB: db}
}

func (s *UserStore) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users(id,email,password)VALUES($1,$2,$3)`
	_, err := s.DB.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	return err
}

func (s *UserStore) FetchUserFromEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id,email,password FROM users WHERE email=$1`
	row := s.DB.QueryRowContext(ctx, query, email)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
