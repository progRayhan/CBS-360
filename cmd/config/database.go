package config

import (
	"fmt"

	"github.com/spf13/viper"
)

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
