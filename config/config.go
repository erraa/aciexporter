package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type APIC struct {
	Username string
	Password string
	BaseURI  string
}

type Log struct {
	Loglevel string
}

type Config struct {
	APIC
	Log
}

func init() {

	viper.SetConfigName("aciexporter")

	// These 2 imports are for testing only
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")

	// This is where you are suppose to have the config file in prod
	viper.AddConfigPath("$HOME")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Could not read config ", err))
	}
}

var ConfigInstance Config

func GetConfig() Config {
	_ = viper.Unmarshal(&ConfigInstance)
	return ConfigInstance
}
