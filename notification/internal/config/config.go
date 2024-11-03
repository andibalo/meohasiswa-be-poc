package config

import (
	"fmt"
	"github.com/andibalo/meowhasiswa-be-poc/notification/pkg/logger"
	"github.com/andibalo/meowhasiswa-be-poc/notification/pkg/trace"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	AppAddress         = ":8082"
	EnvDevEnvironment  = "DEV"
	EnvProdEnvironment = "PROD"
)

type Config interface {
	Logger() *zap.Logger

	AppName() string
	AppEnv() string
	AppAddress() string

	DBConnString() string
	TraceConfig() trace.Config
}

type AppConfig struct {
	logger *zap.Logger
	App    app
	Db     db
	Tracer tracer
}

type app struct {
	AppEnv      string
	AppVersion  string
	Name        string
	Description string
	AppUrl      string
	AppID       string
}

type db struct {
	DSN      string
	User     string
	Password string
	Name     string
	Host     string
	Port     int
	MaxPool  int
}

type tracer struct {
	ServiceName          string
	CollectorURL         string
	CollectorEnvironment string
	Insecure             bool
	FragmentRatio        float64
}

func InitConfig() *AppConfig {
	viper.SetConfigType("env")
	viper.SetConfigName(".env") // name of Config file (without extension)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return &AppConfig{}
	}

	l := logger.GetLogger()

	return &AppConfig{
		logger: l,
		App: app{
			AppEnv:      viper.GetString("APP_ENV"),
			AppVersion:  viper.GetString("APP_VERSION"),
			Name:        "notification-service",
			Description: "notification service",
			AppUrl:      viper.GetString("APP_URL"),
			AppID:       viper.GetString("APP_ID"),
		},
		Db: db{
			DSN:      getRequiredString("DB_DSN"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			Name:     viper.GetString("DB_NAME"),
			MaxPool:  viper.GetInt("DB_MAX_POOLING_CONNECTION"),
		},
		Tracer: tracer{
			ServiceName:          "notification-service",
			CollectorURL:         viper.GetString("OTEL_APM_SERVER_URL"),
			CollectorEnvironment: viper.GetString("OTEL_APM_ENV"),
			Insecure:             viper.GetBool("OTEL_APM_INSECURE"),
			FragmentRatio:        viper.GetFloat64("OTEL_JAEGER_FRACTION_RATIO"),
		},
	}
}

func getRequiredString(key string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}

func (c *AppConfig) Logger() *zap.Logger {
	return c.logger
}

func (c *AppConfig) AppName() string {
	return c.App.Name
}

func (c *AppConfig) AppEnv() string {
	return c.App.AppEnv
}

func (c *AppConfig) AppAddress() string {
	return AppAddress
}

func (c *AppConfig) DBConnString() string {
	return c.Db.DSN
}

func (c *AppConfig) TraceConfig() trace.Config {
	return trace.Config{
		ServiceName:          c.Tracer.ServiceName,
		CollectorURL:         c.Tracer.CollectorURL,
		CollectorEnvironment: c.Tracer.CollectorEnvironment,
		Insecure:             c.Tracer.Insecure,
		FragmentRatio:        c.Tracer.FragmentRatio,
	}
}
