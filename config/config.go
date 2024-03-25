package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	AppVersion string        `json:"appVersion"`
	Host       string        `json:"host" validate:"required"`
	Port       string        `json:"port" validate:"required"`
	N          time.Duration `json:"N" validate:"required"`
	K          int64         `json:"K" validate:"required"`
	Number     int64         `json:"number" validate:"required"`
	Timeout    time.Duration
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DBName   int    `json:"DB"`
}

func LoadConfig() (*viper.Viper, error) {

	viperInstance := viper.New()

	viperInstance.AddConfigPath("./config")
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("yml")

	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viperInstance, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
