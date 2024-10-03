package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Mode     string `mapstructure:"mode"`
	Version  int    `mapstructure:"version"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"database"`
	Server struct {
		Port    string `mapstructure:"port"`
		Host    string `mapstructure:"host"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"server"`
	Services struct {
		Stripe struct {
			Mode       string `mapstructure:"mode"`
			Key        string `mapstructure:"key"`
			SuccessURL string `mapstructure:"success_url"`
			CancelURL  string `mapstructure:"cancel_url"`
		} `mapstructure:"stripe"`
		AMQP struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
		} `mapstructure:"amqp"`
	} `mapstructure:"services"`
}

func New() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "prod"
	}
	viper.SetConfigName(env + ".config")
	viper.SetConfigType("yaml")

	_, caller, _, _ := runtime.Caller(1)

	viper.AddConfigPath(filepath.Join(filepath.Dir(caller), "..", "config"))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var c = &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}
	return c, nil

}
