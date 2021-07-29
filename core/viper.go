package core

import (
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config/app.yml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return v
}
