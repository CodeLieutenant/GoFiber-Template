package logging

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

func Parse(level string) zerolog.Level {
	switch level {
	case "panic":
		return zerolog.PanicLevel
	case "fatal":
		return zerolog.FatalLevel
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	case "info":
		return zerolog.InfoLevel
	}

	return zerolog.Disabled
}

func New(level string, logToConsole bool, writer io.Writer) zerolog.Logger {
	writers := make([]io.Writer, 0, 2)

	if logToConsole {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	writers = append(writers, diode.NewWriter(writer, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages\n", missed)
	}))

	logger := zerolog.New(zerolog.MultiLevelWriter(writers...)).With().Logger()
	logger.Level(Parse(level))

	return logger
}
