package logging

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
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
	}

	return zerolog.InfoLevel
}

func New(level string, writers ...io.Writer) zerolog.Logger {
	if viper.GetBool("logging.console") {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	logger := zerolog.New(io.MultiWriter(writers...)).With().Logger()
	logger.Level(Parse(level))

	return logger
}
