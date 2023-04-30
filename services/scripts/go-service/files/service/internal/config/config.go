package config

import (
	"shared/cutils"

	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
    PORT          string `mapstructure:"PORT"`
    APP_ENV string `mapstructure:"APP_ENV"`
}

func LoadConfig() (c Config, err error) {
    viper.AddConfigPath("./pkg/config/envs")
	if cutils.CheckIsProd() {
        viper.SetConfigName("prod")
        viper.SetConfigType("env")
        } else {
		viper.SetConfigName("dev")
		viper.SetConfigType("env")
	}

    viper.AutomaticEnv()

    err = viper.ReadInConfig()

    if err != nil {
        return
    }

    err = viper.Unmarshal(&c)

    return
}