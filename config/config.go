package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configInstance *Config

type Config struct {
}

func GetConfig() (config Config) {
	if configInstance == nil {
		envFile := os.Getenv("ENV_FILE")
		if envFile == "" {
			envFile = ".env"
		}
		err := godotenv.Load(envFile)
		if err != nil {
			err := godotenv.Load("../../.env") //during testing
			if err != nil {
				panic("Error loading .env file")
			}
		}
		configInstance = new(Config)
	}
	return *configInstance
}

type DatabaseConfig struct {
	DatabaseName string `json:"database_name"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func (Config) getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (c Config) GetEnvVariable(key string) string {
	return c.getEnv(key, "")
}

func (c Config) GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DatabaseName: c.getEnv("POSTGRES_DB_NAME", "test"),
		Host:         c.getEnv("POSTGRES_HOST", "localhost"),
		Port:         c.getEnv("POSTGRES_PORT", "5432"),
		Username:     c.getEnv("POSTGRES_DB_USERNAME", "test"),
		Password:     c.getEnv("POSTGRES_DB_PASSWORD", "test"),
	}
}

func (c Config) GetJwtSecretKey() string {
	return c.getEnv("JWT_SECRET", "aaaa###")
}

func (c Config) GetServerPort() string {
	return c.getEnv("PORT", "8080")
}

func (c Config) GetSchedulerInterval() int {
	value := c.getEnv("SCHEDULER_INTERVAL", "5")
	numb, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error converting string to int, sedning default value", err)
		return 5
	}
	return numb
}
