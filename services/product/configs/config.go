package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Service struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Config struct {
	Environment string `mapstructure:"env"`
	Storage     struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"storage"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"db"`
	} `mapstructure:"database"`
	GrpcServer struct {
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		Timeout uint32 `mapstructure:"timeout"`
	} `mapstructure:"grpc_server"`
	Paginator struct {
		Limit int `mapstructure:"limit"`
	} `mapstructure:"paginator"`
	FileServer struct {
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		Timeout uint32 `mapstructure:"timeout"`
	} `mapstructure:"file_server"`
	Slack struct {
		WebHookUrl string `mapstructure:"webhook_url"`
		Channel    string `mapstructure:"channel"`
		UserName   string `mapstructure:"username"`
	} `mapstructure:"slack"`
	Services struct {
		Admin struct {
			Host string `mapstructure:"host"`
			Port string `mapstructure:"port"`
		} `mapstructure:"admin"`
	} `mapstructure:"services"`
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
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}
	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}
	return c, nil
}
