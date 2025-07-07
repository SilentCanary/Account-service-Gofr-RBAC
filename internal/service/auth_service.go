package service

import (
	"context"
	"errors"

	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/store"
	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/utils"
)

type AuthService struct {
	Store *store.UserStore
}

func NewAuthService(store *store.UserStore) *AuthService {
	return &AuthService{Store: store}
}

func (s *AuthService) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.Store.FetchUserFromEmail(ctx, email)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !utils.MatchPassword(password, user.Password) {
		return "", errors.New("invalid password")
	}

	roles := []string{"reader"}
	token, err := utils.GenerateJWT(user.ID, roles)
	if err != nil {
		return "", err
	}
	return token, nil
}
