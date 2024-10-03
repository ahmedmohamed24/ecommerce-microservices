package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/configs"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/services"
	"github.com/rs/zerolog"
)

func New(c *configs.Config) *slog.Logger {
	slack := services.NewSlackClient(c)
	writers := io.MultiWriter(os.Stdout, slack)
	logger := slog.New(slog.NewJSONHandler(writers, nil))
	return logger
}

type ZeroLoggerWrapper struct {
	Logger *zerolog.Logger
}

func (ZeroLoggerWrapper *ZeroLoggerWrapper) Write(p []byte) (int, error) {
	ZeroLoggerWrapper.Logger.Info().Msg(string(p))
	return len(p), nil
}
