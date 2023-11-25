package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type AppConfig struct {
	Environment string `mapstructure:"environment"`
	Port        int    `mapstructure:"port"`
	Log         Log    `mapstructure:"log"`
	Db          Db     `mapstructure:"db"`
}

type Log struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

type Db struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func GetConfig(configFilePath string) (AppConfig, error) {
	log.Printf("Loading config file: %s", configFilePath)

	conf := viper.New()
	conf.SetConfigFile(configFilePath)

	replacer := strings.NewReplacer(".", "_")
	conf.SetEnvKeyReplacer(replacer)
	conf.AutomaticEnv()

	err := conf.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file at the %s path. Error: %v", configFilePath, err)
	}

	var config AppConfig

	err = conf.Unmarshal(&config)
	if err != nil {
		log.Fatalf("configuration unmarshalling failed!. Error: %v", err)
		return config, err
	}

	return config, nil
}
