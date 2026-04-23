package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Build DSN
func getDsn(cfg *Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.SSLMode,
		cfg.TimeZone,
	)
}

// GORM log level config
func getGormLogLevel() logger.LogLevel {
	if viper.GetBool("DB_DEBUG") {
		return logger.Info
	}
	return logger.Silent
}

// Configure connection pool
func setConnectionPool(db *sql.DB, cfg *Config) {
	maxIdle := cfg.MaxIdleConns
	if maxIdle == 0 {
		maxIdle = 10
	}
	db.SetMaxIdleConns(maxIdle)

	maxOpen := cfg.MaxOpenConns
	if maxOpen == 0 {
		maxOpen = 100
	}
	db.SetMaxOpenConns(maxOpen)

	maxIdleTime := cfg.MaxIdleTime
	if maxIdleTime == 0 {
		maxIdleTime = 5 * time.Minute
	}
	db.SetConnMaxIdleTime(maxIdleTime)

	maxLifetime := cfg.MaxLifeTime
	if maxLifetime == 0 {
		maxLifetime = 30 * time.Minute
	}
	db.SetConnMaxLifetime(maxLifetime)
}

// Main DB connection function
func ConnectDB(cfg *Config) (*gorm.DB, error) {
	dsn := getDsn(cfg)

	maxRetries := cfg.DBMaxRetries
	if maxRetries == 0 {
		maxRetries = 10
	}

	initialDelay := cfg.RetryDelay
	if initialDelay == 0 {
		initialDelay = 2 * time.Second
	}

	delay := initialDelay
	maxDelay := 30 * time.Second
	backoffFactor := 2.0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var dbConn *gorm.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(getGormLogLevel()),
		})

		if err == nil {
			sqlDB, dbErr := dbConn.DB()
			if dbErr == nil {

				// Ping with timeout
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				pingErr := sqlDB.PingContext(ctx)
				cancel()

				if pingErr == nil {
					// Setup pool
					setConnectionPool(sqlDB, cfg)

					if attempt > 1 {
						log.Printf("Connected to database after %d attempts", attempt)
					}

					return dbConn, nil
				}

				// Close bad connection
				_ = sqlDB.Close()
				err = pingErr

			} else {
				err = dbErr
			}
		}

		log.Printf("Database connection attempt %d failed: %v", attempt, err)

		if attempt < maxRetries {
			jitter := time.Duration(r.Int63n(int64(delay)))
			sleep := delay + jitter
			if sleep > maxDelay {
				sleep = maxDelay
			}

			time.Sleep(sleep)

			delay = time.Duration(float64(delay) * backoffFactor)
			if delay > maxDelay {
				delay = maxDelay
			}
		}
	}

	return nil, fmt.Errorf(
		"could not connect to database after %d attempts: %w",
		maxRetries, err,
	)
}
