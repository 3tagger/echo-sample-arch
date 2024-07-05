package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Primary PostgreSQLConfig
}

type PostgreSQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type ServerConfig struct {
	Host string
	Post string
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		Server:   loadServerConfig(),
		Database: loadDatabaseConfig(),
	}, nil
}

func loadServerConfig() ServerConfig {
	return ServerConfig{
		Host: os.Getenv("SERVER_HOST"),
		Post: os.Getenv("SERVER_PORT"),
	}
}

func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Primary: PostgreSQLConfig{
			Host:     os.Getenv("POSTGRESQL_HOST"),
			Port:     os.Getenv("POSTGRESQL_PORT"),
			Username: os.Getenv("POSTGRESQL_USERNAME"),
			Password: os.Getenv("POSTGRESQL_PASSWORD"),
			Name:     os.Getenv("POSTGRESQL_DBNAME"),
		},
	}
}
