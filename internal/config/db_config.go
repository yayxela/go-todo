package config

import "fmt"

type DBConfig struct {
	DBName       string `mapstructure:"DB_NAME"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBAuthSource string `mapstructure:"DB_AUTH_SOURCE"`
}

func (c DBConfig) GetConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBAuthSource,
	)
}
