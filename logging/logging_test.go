package logging_test

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/logging"
)

func TestParse(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	data := []struct {
		level        string
		zerologLevel zerolog.Level
	}{
		{
			level:        "panic",
			zerologLevel: zerolog.PanicLevel,
		},
		{
			level:        "fatal",
			zerologLevel: zerolog.FatalLevel,
		}, {
			level:        "error",
			zerologLevel: zerolog.ErrorLevel,
		}, {
			level:        "warn",
			zerologLevel: zerolog.WarnLevel,
		}, {
			level:        "debug",
			zerologLevel: zerolog.DebugLevel,
		}, {
			level:        "trace",
			zerologLevel: zerolog.TraceLevel,
		}, {
			level:        "info",
			zerologLevel: zerolog.InfoLevel,
		}, {
			level:        "disabled",
			zerologLevel: zerolog.Disabled,
		}, {
			level:        "anything else",
			zerologLevel: zerolog.Disabled,
		},
	}

	for _, lvl := range data {
		assert.Equal(lvl.zerologLevel, logging.Parse(lvl.level))
	}
}

func TestNew(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("WithoutConsole", func(t *testing.T) {
		logger := logging.New("info", false, os.Stderr)
		assert.NotZero(logger)
	})

	t.Run("WithConsole", func(t *testing.T) {
		logger := logging.New("info", true, os.Stderr)
		assert.NotZero(logger)
	})
}
