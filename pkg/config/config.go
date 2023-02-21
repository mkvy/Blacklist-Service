package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

type Config struct {
	Database struct {
		Username   string `env:"PG_USERNAME" env-default:"postgres"`
		Password   string `env:"PG_PASSWORD" env-default:"postgres"`
		DBname     string `env:"DATABASE_NAME" env-default:"blacklist_dev"`
		DriverName string `env:"DRIVER_NAME" env-default:"postgres"`
		HostPort   string `env:"DATABASE_HOST_PORT" env-default:"database:5432"`
	}
	HttpServer struct {
		Host string `env:"SERVER_HOST" env-default:"localhost"`
		Port string `env:"SERVER_PORT" env-default:"8283"`
	}
	Auth struct {
		Username  string `env:"ADMIN_USERNAME" env-default:"admin"`
		Password  string `env:"ADMIN_PASSWORD" env-default:"admin"`
		JwtSecret string `env:"JWT_SECRET" env-default:"secret"`
	}
}

var instance *Config

func GetConfig() *Config {
	once := sync.Once{}
	once.Do(func() {
		instance = &Config{}

		if err := godotenv.Load(".env"); err != nil {
			log.Printf("error during loading environment variables: %s\n", err.Error())
			return
		}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Printf("error during mapping environment variables: %s\n", help)
			return
		}
	})

	return instance
}
