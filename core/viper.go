package core

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config/app.yml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(errors.New("读取配置文件失败：" + err.Error()))
	}

	// 映射到结构体
	if err := v.Unmarshal(&Config); err != nil {
		panic(errors.New("配置映射失败：" + err.Error()))
	}

	return v
}
