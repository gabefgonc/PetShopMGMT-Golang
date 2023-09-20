package config

import (
	"log"
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	App struct {
		Secret string `env:"APP_SECRET"`
	}

	Database struct {
		Name     string `env:"DATABASE_NAME"`
		Host     string `env:"DATABASE_HOST"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		Port     int16  `env:"DATABASE_PORT"`
	}

	Server struct {
		Prefork bool  `env:"PREFORK"`
		Port    int16 `env:"PORT"`
	}
}

func NewConfig() *Config {
	config := &Config{}

	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("Error opening .env file: %s\n", err)
	}

	err = dotenv.NewDecoder(file).Decode(config)

	if err != nil {
		log.Fatalf("Error decoding .env file: %s\n", err)
	}

	return config
}
