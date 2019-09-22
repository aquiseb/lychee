package core

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("global")
	viper.AddConfigPath("_config")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	log.Printf("Config initialized")
}

type Config struct {
	AppName string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret   string
	JWTExpireIn time.Duration

	DebugMode bool
	LogFormat string
}

func LoadConfig() *Config {
	return &Config{
		AppName: viper.Get("app-name").(string),

		DBHost:     viper.Get("db.host").(string),
		DBPort:     viper.Get("db.port").(string),
		DBUser:     viper.Get("db.user").(string),
		DBPassword: viper.Get("db.password").(string),
		DBName:     viper.Get("db.dbname").(string),

		JWTSecret:   viper.Get("auth.jwt-secret").(string),
		JWTExpireIn: viper.GetDuration("auth.jwt-expire-in"),

		DebugMode: viper.Get("log.debug-mode").(bool),
		LogFormat: viper.Get("log.log-format").(string),
	}
}
