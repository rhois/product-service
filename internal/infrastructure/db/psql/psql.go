package db

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // define postgres dialect
)

// NewPsqlRepository initiate postgres database client connection
func NewPsqlRepository(url string) *gorm.DB {
	client, err := gorm.Open("postgres", url)
	if err != nil {
		panic("failed to connect database")
	}
	client.SingularTable(true)
	// Reference: https://www.alexedwards.net/blog/configuring-sqldb
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNECTION"))
	client.DB().SetMaxIdleConns(maxIdleConn)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	maxOpenConn, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_OPEN_CONNECTION"))
	client.DB().SetMaxOpenConns(maxOpenConn)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	connMaxLifeTimeInMinutes, _ := strconv.Atoi(os.Getenv("DATABASE_CONNECTION_MAX_LIFETIME_IN_MINUTE"))
	client.DB().SetConnMaxLifetime(time.Duration(connMaxLifeTimeInMinutes) * time.Minute)

	// show query if log level is debug
	logLevel := os.Getenv("LOG_LEVEL")
	if strings.ToLower(logLevel) == "debug" {
		client.LogMode(true)
	}
	return client
}
