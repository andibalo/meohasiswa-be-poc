package service

import (
	"context"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/request"
)

type authService struct {
	cfg config.Config
}

func NewAuthService(cfg config.Config) AuthService {

	return &authService{
		cfg: cfg,
	}
}

func (s *userService) Register(ctx context.Context, req request.RegisterUserReq) error {
	//ctx, endFunc := trace.Start(ctx, "AuthService.Register", "service")
	//defer endFunc()

	return nil
}
