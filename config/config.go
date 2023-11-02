package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App   App
	Mysql Mysql
}

type App struct {
	TotalWorker int
	CSVFile     string
}

type Mysql struct {
	Host              string
	Port              string
	User              string
	Password          string
	DbName            string
	MaxIdleConnection int
	MaxOpenConnection int
}

func InitConfig() (*Config, error) {
	getwd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "os.Getwd")
	}

	configPath := fmt.Sprintf("%s/config/config.yaml", getwd)

	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
