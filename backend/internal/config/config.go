package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"prod" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTP        HTTP       `yaml:"http_server"`
	DB_Context  DB_Context `yaml:"db_context"`
}

type HTTP struct {
	Address      string        `yaml:"address" env:"HTTP_ADDRESS" env-default:"localhost:8080" env-required:"true"`
	Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
	IddleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

type DB_Context struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DbName   string `yaml:"dbname" env-required:"true"`
}

func MustLoad() *Config {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "/Users/sypher/Desktop/qw/Pet_Project_1-QuikView-/backend/config/config.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config file %v", err)
	}
	return &cfg
}


