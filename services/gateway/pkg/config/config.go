package config

import (
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/utils"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
    Port          string `mapstructure:"PORT"`
    AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
    UsersSvcUrl string `mapstructure:"USERS_SVC_URL"`
    APP_ENV string `mapstructure:"APP_ENV"`
}

func LoadConfig() (c Config, err error) {
    viper.AddConfigPath("./pkg/config/envs")
	if utils.CheckIsProd() {
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