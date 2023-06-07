package config

import (
	"os"
	"path/filepath"

	"github.com/Bulut-Bilisimciler/go-shift-service/logger"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	App struct {
		Env     string `mapstructure:"env"`
		Port    string `mapstructure:"port"`
		Version string `mapstructure:"version"`
	}

	Auth struct {
		JwtPub string `mapstructure:"jwt_pub"`
	}

	DB struct {
		Url string `mapstructure:"url"`
	}

	Cache struct {
		Url string `mapstructure:"url"`
	}

	Broker struct {
		Url           string `mapstructure:"url"`
		ConsumerGroup string `mapstructure:"consumer_group"`
		Topic         string `mapstructure:"topic"`
	}

	Cdn struct {
		Endpoint  string `mapstructure:"endpoint"`
		Region    string `mapstructure:"region"`
		Bucket    string `mapstructure:"bucket"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	}
}

var C config

func ReadConfig(processCwdir string) {
	Config := &C
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(processCwdir, "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.ERROR.Println(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.ERROR.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
