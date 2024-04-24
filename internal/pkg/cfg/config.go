package cfg

import (
	"github.com/spf13/viper"
	"log/slog"
)

type Config struct {
	Port          string
	QueryInterval string
}

const (
	PortKey = "PORT"
)

func NewConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Warn("unable to read the configuration file: " + err.Error())
	}

	viper.AutomaticEnv()

	return &Config{
		Port: viper.GetString(PortKey),
	}
}
