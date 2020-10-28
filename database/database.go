package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Config struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint16
	SslMode  bool
	TimeZone string
}

func (c Config) String() string {
	sslMode := "disable"

	if c.SslMode {
		sslMode = "enable"
	}

	return fmt.Sprintf(
		"user=%s host=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.User, c.Host,
		c.Password, c.DbName, c.Port, sslMode, c.TimeZone)
}

func ConnectDB(c Config) (_ *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(c.String()), &gorm.Config{})

	return db, err
}
