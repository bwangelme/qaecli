package config

import (
	"github.com/spf13/viper"
)

var (
	Server string
)

func InitConfig() {
	Server = viper.GetString("server")
}
