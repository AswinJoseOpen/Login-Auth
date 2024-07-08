package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DbSource string `mapstructure:"DATABASE_SOURCE"`
	Secret   string `mapstructure:"SECRET"`
}

func LoadConfig() (config AppConfig, err error) {
	path := "/Users/aswin.jose/test"
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")
	// viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Unable to load config:%v", err)
	}
	_ = viper.Unmarshal(&config)
	return
}

// func LoadEnv() {
// 	err := godotenv.Load("app.env")
// 	if err != nil {
// 		log.Fatalf("Error loading env files:%v", err)
// 	}
// }
