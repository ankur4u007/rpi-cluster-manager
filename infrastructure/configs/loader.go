package configs

import (
	"fmt"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../configs/")
	err := v.ReadInConfig()
	v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs/")
	err = v.ReadInConfig()
	if err == nil {
		viper.MergeConfigMap(v.AllSettings())
	}
	v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err = v.ReadInConfig()
	if err == nil {
		viper.MergeConfigMap(v.AllSettings())
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("DIETPI")
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&domain.Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
