package config

import (
	"ShortURL/internal/logging"
	"github.com/spf13/viper"
	"sync"
)

func init() {
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
}

var once sync.Once

func GetConfigs() {
	once.Do(func() {
		logger := logging.GetLogger()
		err := viper.ReadInConfig()
		if err != nil {
			logger.Fatal(err)
		}
	})
}
