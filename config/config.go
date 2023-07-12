package config

import (
	"os"
	"path/filepath"

	"buluttan/shift-service/logger"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	App struct {
		Env     string `mapstructure:"env"`
		Port    string `mapstructure:"port"`
		Version string `mapstructure:"version"`
	} `mapstructure:"app"`

	Auth struct {
		JwtPub string `mapstructure:"jwt_pub"`
	} `mapstructure:"auth"`

	DB struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"db"`

	Cache struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"cache"`

	Broker struct {
		Url           string `mapstructure:"url"`
		ConsumerGroup string `mapstructure:"consumer_group"`
		Topic         string `mapstructure:"topic"`
	} `mapstructure:"broker"`

	Cdn struct {
		Endpoint  string `mapstructure:"endpoint"`
		Region    string `mapstructure:"region"`
		Bucket    string `mapstructure:"bucket"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"cdn"`
}

var C config

func ReadConfig(processCwdir string) {
	Config := &C
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(processCwdir, "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.ERROR.Println("Cannot read config file:", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.ERROR.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
