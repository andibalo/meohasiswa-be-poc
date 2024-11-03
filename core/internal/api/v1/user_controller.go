package v1

import (
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/request"
	"github.com/andibalo/meowhasiswa-be-poc/core/internal/service"
	"github.com/andibalo/meowhasiswa-be-poc/core/pkg/trace"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserController struct {
	cfg     config.Config
	userSvc service.UserService
}

func NewUserController(cfg config.Config, userSvc service.UserService) *UserController {

	return &UserController{
		cfg:     cfg,
		userSvc: userSvc,
	}
}

func (h *UserController) AddRoutes(r *gin.Engine) {
	ur := r.Group("/api/v1/user")

	ur.GET("/test", h.TestLog)
	ur.POST("/test", h.TestLogWithBody)
	ur.POST("/test-call-notif", h.TestCallNotifService)
}

func (h *UserController) TestLog(c *gin.Context) {
	_, endFunc := trace.Start(c.Copy().Request.Context(), "UserController.TestLog", "controller")
	defer endFunc()

	h.cfg.Logger().Info("test log from core service")

	c.JSON(http.StatusOK, nil)
}

func (h *UserController) TestLogWithBody(c *gin.Context) {
	_, endFunc := trace.Start(c.Copy().Request.Context(), "UserController.TestLogWithBody", "controller")
	defer endFunc()

	var data request.TestLogWithBodyReq

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	h.cfg.Logger().Info("test log from core service", zap.String("reqmsg", data.Msg))

	c.JSON(http.StatusOK, nil)
}

func (h *UserController) TestCallNotifService(c *gin.Context) {
	ctx, endFunc := trace.Start(c.Copy().Request.Context(), "UserController.TestCallNotifService", "controller")
	defer endFunc()

	var data request.TestCallNotifServiceReq

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	h.cfg.Logger().Info("test call notif service from core service", zap.String("reqmsg", data.TemplateName))

	err := h.userSvc.TestCallNotifService(ctx, data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
