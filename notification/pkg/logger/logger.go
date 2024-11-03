package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	initLoggerOnce sync.Once
	logger         *zap.Logger
)

func GetLogger() *zap.Logger {
	initLoggerOnce.Do(func() {
		logger = InitLogger()

	})

	return logger
}

func InitLogger() *zap.Logger {

	var zapLogger *zap.Logger

	config := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stdout"}, // Log to file
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			MessageKey:     "msg",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
	}

	zapLogger, err := config.Build()

	if err != nil {
		panic(fmt.Sprintf("logger initialization failed %v", err))
	}

	if os.Getenv("APP_ENV") == "DEV" {
		zapLogger, err = zap.NewDevelopment()

		if err != nil {
			panic(fmt.Sprintf("logger initialization failed %v", err))
		}
	}

	zapLogger.Info("logger started")

	defer zapLogger.Sync()

	return zapLogger
}
