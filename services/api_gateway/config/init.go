package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	AdminService struct {
		Host string `mapstructure:"grpc_host"`
		Port string `mapstructure:"grpc_port"`
	} `mapstructure:"admins"`
	ApiGateway struct {
		Host string `mapstructure:"grpc_host"`
		Port string `mapstructure:"grpc_port"`
	} `mapstructure:"api_gateway"`
	CustomerService struct {
		Host string `mapstructure:"grpc_host"`
		Port string `mapstructure:"grpc_port"`
	} `mapstructure:"customers"`
	OrderService struct {
		Host string `mapstructure:"grpc_host"`
		Port string `mapstructure:"grpc_port"`
	} `mapstructure:"orders"`
	PaymentService struct {
		Host string `mapstructure:"grpc_host"`
		Port string `mapstructure:"grpc_port"`
	} `mapstructure:"payments"`
	ProductService struct {
		Host string `mapstructure:"grpc_host"`
		Port string `mapstructure:"grpc_port"`
	} `mapstructure:"products"`
}

func New() (*Config, error) {
	var cfg Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	_, c, _, _ := runtime.Caller(0)
	path := filepath.Dir(c)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
