package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Key string `mapstructure:"key"`
	} `mapstructure:"app"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"database"`
	} `mapstructure:"database"`

	GRPC struct {
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"grpc"`
}

func New() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "prod"
	}
	viper.SetConfigName(env + ".config")
	viper.SetConfigType("yaml")
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	viper.AddConfigPath(basepath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %w \n", err)
	}
	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("Fatal error while trying to unmarshal config: %w \n", err)
	}
	return config, nil

}
