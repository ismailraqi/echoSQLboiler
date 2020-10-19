package routers

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "admin",
		"password": "password1",
		"host":     "localhost",
		"port":     3306,
		"database": "test",
	}
	configName  = "config"
	configPaths = []string{
		".",
	}
)

//Configuration config func
func Configuration() Config {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(configName)
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not decode config into struct: %v", err)
	}
	fmt.Printf("Username from struct: %s\n", config.Username)
	fmt.Printf("Config struct: %v\n", config)
	return config
}
