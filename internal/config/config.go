package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig     AppConfig     `mapstructure:",squash"`
	DBConfig      DBConfig      `mapstructure:",squash"`
	SwaggerConfig SwaggerConfig `yaml:"swagger"`
}

// New ...
// Загружает конфиги в структуру Config
func New() (*Config, error) {
	var config Config
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	viper.AddConfigPath("./api/swagger/swagger-config.yaml")
	if err := viper.MergeInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
