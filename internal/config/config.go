package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	// Database
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	SSLMode    string
	TimeZone   string

	// Pool settings
	MaxIdleConns int
	MaxOpenConns int
	MaxIdleTime  time.Duration
	MaxLifeTime  time.Duration

	// Retry
	DBMaxRetries int
	RetryDelay   time.Duration

	// Debug
	Debug bool
}

var AppConfig *Config

func Load() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	AppConfig = &Config{
		DBName:     viper.GetString("MASTER_DB_NAME"),
		DBUser:     viper.GetString("MASTER_DB_USER"),
		DBPassword: viper.GetString("MASTER_DB_PASSWORD"),
		DBHost:     viper.GetString("MASTER_DB_HOST"),
		DBPort:     viper.GetString("MASTER_DB_PORT"),
		SSLMode:    viper.GetString("MASTER_SSL_MODE"),
		TimeZone:   viper.GetString("TIME_ZONE"),

		MaxIdleConns: viper.GetInt("SET_MAX_IDLE_CONNECTIONS"),
		MaxOpenConns: viper.GetInt("SET_MAX_OPEN_CONNECTIONS"),
		MaxIdleTime:  viper.GetDuration("SET_CONNECTION_MAX_IDLE_TIME"),
		MaxLifeTime:  viper.GetDuration("SET_CONNECTION_MAX_LIFE_TIME"),

		DBMaxRetries: viper.GetInt("DB_MAX_RETRIES"),
		RetryDelay:   viper.GetDuration("DB_INITIAL_DELAY"),

		Debug: viper.GetBool("DEBUG"),
	}
}
