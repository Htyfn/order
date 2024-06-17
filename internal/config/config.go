package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type OldConfig struct {
	Env               string
	LogLevel          int
	LogFormat         string
	ServerAddress     string
	ServerTimeout     time.Duration
	ServerIdleTimeout time.Duration
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
}

type Config struct {
	Env             string `yaml:"env" env-default:"local"`
	LogLevel        int    `yaml:"loglevel" env-default:"0"`
	LogFormat       string `yaml:"logformat" env-default:"text"`
	Storage         `yaml:"storage" env-required:"true"`
	HTTPServer      `yaml:"http_server"`
	ProductEndpoint string `yaml:"product-endpoint" env-required:"true"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "C:/UserData/programming/order/config/dev.yml"
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	os.Setenv("STORAGE_CONFIG",
		fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			cfg.Storage.Host,
			cfg.Storage.Port,
			cfg.Storage.User,
			cfg.Storage.Password,
			cfg.Storage.DBName))

	return &cfg
}
