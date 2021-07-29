package core

import (
	"github.com/spf13/viper"
)

var AppConfig *viper.Viper

func init() {

	v := viper.New()
	v.SetConfigFile("config/app.yml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	AppConfig = v
}
