package config

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

var (
	appConfig *config
	once      sync.Once
)

type config struct {
	Port       string // flag and env
	IsDev      bool   // flag
	DBHost     string // env
	DBPort     string // env
	DBName     string // env
	DBUser     string // env
	DBPassword string // env
	SecretKey  string // env
}

func (c *config) GetDataSourceName() string {
	sslmode := "enable"
	if c.IsDev {
		sslmode = "disable"
	}
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBName, c.DBUser, c.DBPassword, sslmode,
	)
}

func NewConfig() *config {
	once.Do(func() {
		var isDev bool
		flag.BoolVar(&isDev, "dev", false, "enable development mode")
		flag.Parse()
		fmt.Println(isDev)

		appConfig = &config{
			Port:       getEnv("PORT", "8081"),
			IsDev:      isDev,
			DBHost:     getEnv("MY_DB_HOST", "localhost"),
			DBPort:     getEnv("MY_DB_PORT", "5433"),
			DBName:     getEnv("MY_DB_NAME", "sandbox"),
			DBUser:     getEnv("MY_DB_USER", "postgres"),
			DBPassword: getEnv("MY_DB_PASSWORD", "postgres"),
			SecretKey:  getEnv("MY_SECRET_KEY", "12345678901234567890123456789012"),
		}
	})
	return appConfig
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
