package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/gommon/log"
)

type (
	Config struct {
		Env        string `yaml:"env" env-default:"local"`
		HTTPServer `yaml:"http_server"`
		DB         `yaml:"db"`
	}

	HTTPServer struct {
		HTTPAddress string        `yaml:"address" env-default:"localhost:8080"`
		TimeOut     time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeOut time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}

	DB struct {
		DBAddress string `yaml:"address" env-default:"localhost:5433"`
		User      string `yaml:"user" env-required:"true"`
		Password  string `yaml:"password" env-required:"true"`
		NameDB    string `yaml:"name_db" env-required:"true"`
	}
)

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is no set")
	}

	// check if file exist

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file is not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
