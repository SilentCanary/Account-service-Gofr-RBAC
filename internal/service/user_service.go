package service

import (
	"context"

	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/models"
	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/store"
	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/utils"
	"github.com/google/uuid"
)

type UserService struct {
	Store *store.UserStore
}

func NewUserService(store *store.UserStore) *UserService {
	return &UserService{Store: store}
}

func (s *UserService) RegisterUser(ctx context.Context, email, password string) error {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := &models.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hashed,
	}
	return s.Store.CreateUser(ctx, user)
}
