package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	isDebug := viper.GetBool("APP_DEBUG")
	engine := routes(isDebug)

	port := fmt.Sprintf(":%d", viper.GetInt("PORT"))

	engine.Run(port)
}
