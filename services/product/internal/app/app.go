package app

import (
	"log/slog"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/configs"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/database"
	"github.com/ahmedmohamed24/ecommerce-microservices/product/internal/logger"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Config *configs.Config
	Logger *slog.Logger
}

func New() (*App, error) {
	c, err := configs.New()
	if err != nil {
		return nil, err
	}
	db, err := database.New(c)
	if err != nil {
		return nil, err
	}
	return &App{
		DB:     db,
		Logger: logger.New(c),
		Config: c,
	}, nil
}
