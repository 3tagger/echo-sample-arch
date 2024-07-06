package config

import (
	"os"
	"strconv"

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
	Host        string
	Post        string
	GracePeriod int
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	serverCfg, err := loadServerConfig()
	if err != nil {
		return Config{}, err
	}

	dbCfg := loadDatabaseConfig()

	return Config{
		Server:   serverCfg,
		Database: dbCfg,
	}, nil
}

func loadServerConfig() (ServerConfig, error) {
	num, err := strconv.ParseInt(os.Getenv("SERVER_GRACE_PERIOD"), 10, 32)
	if err != nil {
		return ServerConfig{}, err
	}

	return ServerConfig{
		Host:        os.Getenv("SERVER_HOST"),
		Post:        os.Getenv("SERVER_PORT"),
		GracePeriod: int(num),
	}, nil
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
