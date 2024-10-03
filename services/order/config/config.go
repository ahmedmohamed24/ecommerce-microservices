package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string `mapstructure:"service_name"`
	Server      struct {
		Version string `mapstructure:"version"`
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	Services struct {
		Customer struct {
			Host string `mapstructure:"host"`
			Port string `mapstructure:"port"`
		} `mapstructure:"customer"`
		Product struct {
			Host string `mapstructure:"host"`
			Port string `mapstructure:"port"`
		} `mapstructure:"product"`
		Payment struct {
			Host string `mapstructure:"host"`
			Port string `mapstructure:"port"`
		} `mapstructure:"payment"`
		AMQP struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
		} `mapstructure:"amqp"`
	}
}

func New() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "prod"
	}
	_, c, _, _ := runtime.Caller(0)
	viper.SetConfigName(env + ".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(c))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config = &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
