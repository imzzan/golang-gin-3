package config

import "github.com/spf13/viper"

type Config struct {
	PORT   string
	DB_URL string
	TOKEN  string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		panic(err)
	}
}
