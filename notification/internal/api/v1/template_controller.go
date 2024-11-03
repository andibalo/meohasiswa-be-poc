package v1

import (
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/request"
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/response"
	"github.com/andibalo/meowhasiswa-be-poc/notification/pkg/trace"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type TemplateController struct {
	cfg config.Config
}

func NewTemplateController(cfg config.Config) *TemplateController {

	return &TemplateController{
		cfg: cfg,
	}
}

func (h *TemplateController) AddRoutes(r *gin.Engine) {
	nc := r.Group("/api/v1/template")

	nc.POST("", h.CreateTemplate)
}

func (h *TemplateController) CreateTemplate(c *gin.Context) {
	_, endFunc := trace.Start(c.Copy().Request.Context(), "TemplateController.CreateTemplate", "controller")
	defer endFunc()

	var data request.CreateTemplateReq

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	h.cfg.Logger().Info("test create template", zap.String("reqmsg", data.TemplateName))

	c.JSON(http.StatusOK, response.CreateNotifTemplateResp{Success: true})
}
