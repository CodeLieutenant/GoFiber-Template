package logging

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DateTimeFormat = "2006-01-02 15:04:05"

func ConfigureDefaultLogger(level string) {
	zerologLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		panic("Failed to parse logging level: " + level)
	}

	zerolog.SetGlobalLevel(zerologLevel)
	zerolog.TimeFieldFormat = DateTimeFormat
	zerolog.DurationFieldUnit = time.Microsecond
	zerolog.TimestampFunc = time.Now().UTC

	log.Logger = log.Output(zerolog.NewConsoleWriter())
}

func New(level string) zerolog.Logger {
	var logger zerolog.Logger

	zerologLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		panic("Failed to parse logging level: " + level)
	}

	logger = zerolog.New(zerolog.MultiLevelWriter(zerolog.NewConsoleWriter())).
		With().
		Timestamp().
		Logger().
		Level(zerologLevel)

	return logger
}
