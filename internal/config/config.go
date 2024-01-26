package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

type Config struct {
	Env      string `yaml:"env" env-required:"true"`
	Database struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
		User string `yaml:"user" env-required:"true"`
		Pass string `yaml:"pass" env-required:"true"`
		Name string `yaml:"name" env-required:"true"`
	} `yaml:"database"`
	JWT struct {
		Secret string `yaml:"secret" env-required:"true"`
		TTL    string `yaml:"ttl" env-required:"true"`
	}
	GRPC struct {
		Port int `yaml:"port" env-required:"true"`
	}
	Redis struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
		Pass string `yaml:"pass" env-required:"true"`
		DB   int    `yaml:"db" env-default:"0"`
	} `yaml:"redis"`
}

var once sync.Once
var config Config

func mustLoad() {
	configPath := ""
	if os.Getenv("DOCKER_ENV") == "true" {
		configPath = "./config/config_docker.yaml"
	}
	if os.Getenv("TEST_ENV") == "true" {
		configPath = "./config/config_test.yaml"
	}
	if configPath == "" {
		configPath = "./config/config_local.yaml"
	}
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		panic(err)
	}
}

func Get() Config {
	once.Do(mustLoad)
	return config
}
