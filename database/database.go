package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"time"
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
		"user=%s host=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.User,
		c.Host,
		c.Password,
		c.DbName,
		c.Port,
		sslMode,
		c.TimeZone,
	)
}

func ConnectDB(c fmt.Stringer, w io.Writer) (_ *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(c.String()), &gorm.Config{
		PrepareStmt: true,
		Logger: logger.New(log.New(w, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      true,
		}),
	})

	return db, err
}
