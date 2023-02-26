package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Port             string `mapstructure:"PORT"`
	AccountClientUrl string `mapstructure:"ACCOUNT_CLIENT_URL"`
	ProductClientUrl string `mapstructure:"PRODUCT_CLIENT_URL"`
	OrderClientUrl   string `mapstructure:"ORDER_CLIENT_URL"`
}

var Config Configuration

func (config *Configuration) LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(config)
	return
}
