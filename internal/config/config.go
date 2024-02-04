package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
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
	} `yaml:"jwt"`
	GRPC struct {
		Port int `yaml:"port" env-required:"true"`
	} `yaml:"grpc"`
	Redis struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
		Pass string `yaml:"pass" env-required:"true"`
		DB   int    `yaml:"db" env-default:"0"`
	} `yaml:"redis"`
	Loki struct {
		Host string `yaml:"host" env-required:"true"`
		Port int    `yaml:"port" env-required:"true"`
	} `yaml:"loki"`
	MinIO struct {
		Host      string `yaml:"host" env-required:"true"`
		Port      int    `yaml:"port" env-required:"true"`
		AccessKey string `yaml:"access_key" env-required:"access_key"`
		SecretKey string `yaml:"secret_key" env-required:"true"`
		SSLMode   bool   `yaml:"use_ssl" env-default:"false"`
		Bucket    string `yaml:"bucket" env-required:"true"`
	} `yaml:"minio"`
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

func GetLogger() *slog.Logger {
	conf := Get()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger = logger.With("service", "eshop-products-ms").With("env", conf.Env)

	return logger
}
