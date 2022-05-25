package helpers

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBuser     string `mapstructure:"DB_USER"`
	DBpassword string `mapstructure:"DB_PASSWORD"`
	DBname     string `mapstructure:"DB_NAME"`
	DBhost     string `mapstructure:"DB_HOST"`
	DBport     int    `mapstructure:"DB_PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	//to set different type of env
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
