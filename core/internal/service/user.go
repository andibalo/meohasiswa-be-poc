package service

import (
	"context"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/request"
	"github.com/andibalo/meowhasiswa-be-poc/core/pkg/integration/notifsvc"
	"github.com/andibalo/meowhasiswa-be-poc/core/pkg/trace"
)

type userService struct {
	cfg      config.Config
	notifSvc notifsvc.INotifSvc
}

func NewUserService(cfg config.Config, notifSvc notifsvc.INotifSvc) UserService {

	return &userService{
		cfg:      cfg,
		notifSvc: notifSvc,
	}
}

func (s *userService) TestCallNotifService(ctx context.Context, req request.TestCallNotifServiceReq) error {
	ctx, endFunc := trace.Start(ctx, "UserService.TestCallNotifService", "service")
	defer endFunc()

	_, err := s.notifSvc.CreateNotifTemplate(ctx, notifsvc.CreateNotifTemplateReq{TemplateName: req.TemplateName})
	if err != nil {
		return err
	}

	return nil
}
