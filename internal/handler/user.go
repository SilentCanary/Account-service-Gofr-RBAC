package handler

import (
	"github.com/SilentCanary/Account-service-Gofr-RBAC/internal/service"
	"gofr.dev/pkg/gofr"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) RegisterUser(ctx *gofr.Context) (interface{}, error) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.Bind(&body); err != nil {
		return nil, err
	}
	err := h.Service.RegisterUser(ctx, body.Email, body.Password)
	if err != nil {
		return nil, err
	}
	return map[string]string{"message": "user registered successfully"}, nil
}
