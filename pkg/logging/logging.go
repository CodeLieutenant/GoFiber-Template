package logging

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DateTimeFormat = "2006-01-02 15:04:05"

func ConfigureDefaultLogger(level string, prettyPrint bool) {
	zerologLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		panic("Failed to parse logging level: " + level)
	}

	zerolog.SetGlobalLevel(zerologLevel)
	zerolog.TimeFieldFormat = DateTimeFormat
	zerolog.DurationFieldUnit = time.Microsecond
	zerolog.TimestampFunc = time.Now().UTC

	var w io.Writer

	if prettyPrint {
		w = zerolog.NewConsoleWriter()
	} else {
		w = os.Stdout
	}

	log.Logger = log.Output(w)
}

func New(level string, prettyPrint bool) zerolog.Logger {
	var logger zerolog.Logger

	zerologLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		panic("Failed to parse logging level: " + level)
	}

	var w io.Writer

	if prettyPrint {
		w = zerolog.NewConsoleWriter()
	} else {
		w = os.Stdout
	}

	logger = zerolog.New(w).
		With().
		Timestamp().
		Logger().
		Level(zerologLevel)

	return logger
}
