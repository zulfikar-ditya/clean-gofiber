package config

import "github.com/spf13/viper"

type APP_CONFIG struct {
	APP_NAME string
	APP_ENV  string
	APP_PORT string
	APP_KEY  string
	APP_TIMEZONE string

	CLIENT_URL string
}

var appConfig *APP_CONFIG

func AppConfig(v *viper.Viper) APP_CONFIG {
	cfg := APP_CONFIG{
		APP_NAME: v.GetString("APP_NAME"),
		APP_ENV:  v.GetString("APP_ENV"),
		APP_PORT: v.GetString("APP_PORT"),
		APP_KEY:  v.GetString("APP_KEY"),
		APP_TIMEZONE: v.GetString("APP_TIMEZONE"),

		CLIENT_URL: v.GetString("CLIENT_URL"),
	}

	appConfig = &cfg

	return cfg
}

func Get() *APP_CONFIG {
	return appConfig
}

func Env(key string) string {
	if appConfig == nil {
		return ""
	}

	switch key {
	case "APP_NAME":
		return appConfig.APP_NAME
	case "APP_ENV":
		return appConfig.APP_ENV
	case "APP_PORT":
		return appConfig.APP_PORT
	case "APP_KEY":
		return appConfig.APP_KEY
	case "APP_TIMEZONE":
		return appConfig.APP_TIMEZONE
	case "CLIENT_URL":
		return appConfig.CLIENT_URL
	default:
		return ""
	}
}
