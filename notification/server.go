package notification

import (
	"context"
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/api"
	v1 "github.com/andibalo/meowhasiswa-be-poc/notification/internal/api/v1"
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/notification/pkg/trace"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	gin *gin.Engine
	srv *http.Server
}

func NewServer(cfg config.Config, tracer *trace.Tracer) *Server {

	router := gin.New()

	//router.Use(middleware.RequestLogger())

	tracer.SetGinMiddleware(router, cfg.AppName())

	router.Use(trace.TracerLogger())

	router.Use(gin.Recovery())

	uc := v1.NewUserController(cfg)

	tc := v1.NewTemplateController(cfg)
	
	registerHandlers(router, &api.HealthCheck{}, uc, tc)

	return &Server{
		gin: router,
	}
}

func (s *Server) Start(addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: s.gin,
	}

	s.srv = srv

	return srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {

	return s.srv.Shutdown(ctx)
}

func registerHandlers(g *gin.Engine, handlers ...api.Handler) {
	for _, handler := range handlers {
		handler.AddRoutes(g)
	}
}
