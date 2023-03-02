package helpers

import (
	"github.com/spf13/viper"
	"log"
)

func GetEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err.Error())
	}
	val, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("error when getting env value")
	}

	return val
}
