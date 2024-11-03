package service

import (
	"context"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/request"
)

type UserService interface {
	TestCallNotifService(ctx context.Context, req request.TestCallNotifServiceReq) error
}
