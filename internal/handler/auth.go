package handler

// TODO: will add auth handler later
import (
	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/service"
	"gofr.dev/pkg/gofr"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) LoginUser(ctx *gofr.Context) (interface{}, error) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.Bind(&body); err != nil {
		return nil, err
	}
	token, err := h.Service.LoginUser(ctx, body.Email, body.Password)
	if err != nil {
		return nil, err
	}

	return map[string]string{"token": token}, nil
}
