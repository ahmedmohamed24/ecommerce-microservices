package database

import (
	"fmt"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.Database.Host, c.Database.User, c.Database.Password, c.Database.Name, c.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
