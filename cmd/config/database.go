package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDsn() string {
	dbName := viper.GetString("MASTER_DB_NAME")
	dbUser := viper.GetString("MASTER_DB_USER")
	dbPassword := viper.GetString("MASTER_DB_PASSWORD")
	dbHost := viper.GetString("MASTER_DB_HOST")
	dbPort := viper.GetString("MASTER_DB_PORT")
	dbSSLMode := viper.GetString("MASTER_SSL_MODE")
	dbTimeZone := viper.GetString("TIME_ZONE")

	dsn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s TimeZone=%s",
		dbName, dbUser, dbPassword, dbHost, dbPort, dbSSLMode, dbTimeZone,
	)

	return dsn
}

func retryDatabaseConnection(dsn string, maxRetries int, initialDelay time.Duration) (*gorm.DB, error) {
	var dbConn *gorm.DB
	var err error
	delay := initialDelay
	maxDelay := 30 * time.Second
	backoffFactor := 2.0

	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Try to connect
		dbConn, err = gorm.Open(postgres.Open(dsn))
		if err == nil {
			// Connection successful, test it
			sqlDb, err := dbConn.DB()
			if err == nil {
				err = sqlDb.Ping()
				if err == nil {
					// Configure connection pool
					sqlDb.SetMaxIdleConns(viper.GetInt("SET_MAX_IDLE_CONNECTIONS"))
					sqlDb.SetMaxOpenConns(viper.GetInt("SET_MAX_OPEN_CONNECTIONS"))
					sqlDb.SetConnMaxIdleTime(viper.GetDuration("SET_CONNECTION_MAX_IDLE_TIME"))
					sqlDb.SetConnMaxLifetime(viper.GetDuration("SET_CONNECTION_MAX_LIFE_TIME"))

					if attempt > 1 {
						// set logger
					}
					return dbConn, nil
				}
			}
		}

		if attempt < maxRetries {
			// logger

			time.Sleep(delay)

			delay = time.Duration(float64(delay) * backoffFactor)
			if delay > maxDelay {
				delay = maxDelay
			}
		}
	}

	// logger

	return nil, err
}

func ConnectDB() {
	var err error
	dsn := getDsn()
	// slowSQLThreshold := viper.GetDuration("SLOW_SQL_LOG_THRESHOLD") * time.Millisecond

	maxRetries := 10
	initialDelay := 2 * time.Second

	db, err = retryDatabaseConnection(dsn, maxRetries, initialDelay)
	if err != nil {
		// log
	}
}

func GetDatabase() *gorm.DB {
	return db
}
