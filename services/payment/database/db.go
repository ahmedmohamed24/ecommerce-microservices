package database

import (
	"fmt"

	"github.com/ahmedmohamed24/ecommerce-microservices/payment/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s", c.Database.Host, c.Database.User, c.Database.Password, c.Database.Database, c.Database.Port, "disable", "UTC")
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
