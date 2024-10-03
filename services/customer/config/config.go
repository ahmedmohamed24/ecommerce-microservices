package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string `mapstructure:"service_name"`
	Mode        string `mapstructure:"mode"`
	APP_KEY     string `mapstructure:"app_key"`
	Server      struct {
		Version int    `mapstructure:"version"`
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"database"`
	} `mapstructure:"database"`
}

func New() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "prod"
	}
	_, dir, _, _ := runtime.Caller(0)
	viper.SetConfigFile(filepath.Join(filepath.Dir(dir), env+".config.yaml"))
	c := &Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}
	viper.Unmarshal(c)
	return c, nil
}
