package v1

import (
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/request"
	"github.com/andibalo/meowhasiswa-be-poc/notification/pkg/trace"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type NotificationController struct {
	cfg config.Config
}

func NewUserController(cfg config.Config) *NotificationController {

	return &NotificationController{
		cfg: cfg,
	}
}

func (h *NotificationController) AddRoutes(r *gin.Engine) {
	nc := r.Group("/api/v1/notification")

	nc.GET("/test", h.TestLog)
	nc.POST("/test", h.TestLogWithBody)
}

func (h *NotificationController) TestLog(c *gin.Context) {
	h.cfg.Logger().Info("test log from notification service")

	c.JSON(http.StatusOK, nil)
}

func (h *NotificationController) TestLogWithBody(c *gin.Context) {
	_, endFunc := trace.Start(c.Copy().Request.Context(), "NotificationController.TestLogWithBody", "controller")
	defer endFunc()

	var data request.TestLogWithBodyReq

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	h.cfg.Logger().Info("test log from core service", zap.String("reqmsg", data.Msg))

	c.JSON(http.StatusOK, nil)
}
