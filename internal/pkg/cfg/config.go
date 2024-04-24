package cfg

import (
	"github.com/spf13/viper"
	"log/slog"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBTimeZone string
}

const (
	PortKey       = "PORT"
	DBUserKey     = "DB_USER"
	DBPasswordKey = "DB_PASSWORD"
	DBHostKey     = "DB_HOST"
	DBPortKey     = "DB_PORT"
	DBNameKey     = "DB_NAME"
	DBTimeZoneKey = "DB_TIMEZONE"
)

func NewConfig() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		slog.Warn("unable to read the configuration file: " + err.Error())
	}

	viper.AutomaticEnv()

	return &Config{
		Port:       viper.GetString(PortKey),
		DBUser:     viper.GetString(DBUserKey),
		DBPassword: viper.GetString(DBPasswordKey),
		DBHost:     viper.GetString(DBHostKey),
		DBPort:     viper.GetString(DBPortKey),
		DBName:     viper.GetString(DBNameKey),
		DBTimeZone: viper.GetString(DBTimeZoneKey),
	}
}
