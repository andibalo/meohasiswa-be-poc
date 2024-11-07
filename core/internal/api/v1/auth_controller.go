package v1

import (
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	cfg     config.Config
	authSvc service.AuthService
}

func NewAuthController(cfg config.Config, authSvc service.AuthService) *AuthController {

	return &AuthController{
		cfg:     cfg,
		authSvc: authSvc,
	}
}

func (h *AuthController) AddRoutes(r *gin.Engine) {
	// ur := r.Group("/api/v1/user")

}
