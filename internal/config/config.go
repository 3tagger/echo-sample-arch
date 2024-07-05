package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host string
	Post string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return Config{
		Server: loadServerConfig(),
	}

}

func loadServerConfig() ServerConfig {
	return ServerConfig{
		Host: os.Getenv("SERVER_HOST"),
		Post: os.Getenv("SERVER_PORT"),
	}
}
