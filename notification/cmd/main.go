package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/andibalo/meowhasiswa-be-poc/notification"
	"github.com/andibalo/meowhasiswa-be-poc/notification/internal/config"
	"github.com/andibalo/meowhasiswa-be-poc/notification/pkg/trace"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.InitConfig()

	//database := db.InitDB(cfg)

	tracer := initTracer(cfg)

	server := notification.NewServer(cfg, tracer)

	cfg.Logger().Info(fmt.Sprintf("Server starting at port %s", cfg.AppAddress()))

	go func() {
		if err := server.Start(cfg.AppAddress()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			cfg.Logger().Fatal("failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cfg.Logger().Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		cfg.Logger().Fatal("Server force to shutdown")
	}

	cfg.Logger().Info("Server exiting")
}

func initTracer(cfg config.Config) *trace.Tracer {

	traceConfig := cfg.TraceConfig()

	// init tracer type
	tracer, err := trace.Init(context.Background(), traceConfig)
	if err != nil {
		log.Fatal("error init tracer: ", err)
	}

	return tracer
}
