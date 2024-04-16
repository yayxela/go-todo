package config

type AppConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	BasePath string `mapstructure:"BASE_PATH"`
	Timezone string `mapstructure:"TIMEZONE"`
}
