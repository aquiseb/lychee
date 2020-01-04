package core

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// InitViper initializes viper configs
func InitViper() {
	viper.SetConfigName("global")
	viper.AddConfigPath("_config")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	fmt.Println("Config initialized")
}
