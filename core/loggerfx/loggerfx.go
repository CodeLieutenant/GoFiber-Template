package loggerfx

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/nano-interactive/go-utils"
	appLogger "github.com/nano-interactive/go-utils/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"go.uber.org/fx"
)

type (
	SinkType int
	Sink     struct {
		Level       string
		Args        []any
		Type        SinkType
		PrettyPrint bool
	}
)

const (
	Stdout SinkType = iota
	Stderr
	File
)

var (
	ErrArgForFileNotProvided    = errors.New("output file path must be provided for SINK FILE")
	ErrUnexpectedArgForFileType = errors.New("invalid type for the sink FILE => expected string or fmt.Stringer")
	ErrInvalidSinkType          = errors.New("invalid sink type")
)

func Module(sink Sink) fx.Option {
	return fx.Module("ZerologLogger", fx.Provide(
		func(lc fx.Lifecycle) (zerolog.Logger, error) {
			w, closer, err := getWriter(sink)
			if err != nil {
				return zerolog.Logger{}, err
			}

			if closer != nil {
				lc.Append(fx.StopHook(closer))
			}

			return appLogger.New(sink.Level, sink.PrettyPrint).
				Output(w).
				With().
				Stack().
				Logger(), nil
		}),
		fx.Invoke(func(lc fx.Lifecycle) error {
			w, closer, err := getWriter(sink)
			if err != nil {
				return err
			}

			if closer != nil {
				lc.Append(fx.StopHook(closer))
			}

			//nolint:all
			zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
			appLogger.ConfigureDefaultLogger(sink.Level, sink.PrettyPrint, w)

			return nil
		}),
	)
}

func getWriter(sink Sink) (io.Writer, func() error, error) {
	switch sink.Type {
	case Stdout:
		if sink.PrettyPrint {
			return zerolog.NewConsoleWriter(), nil, nil
		}

		return os.Stdout, nil, nil
	case Stderr:
		if sink.PrettyPrint {
			return zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
				w.Out = os.Stderr
			}), nil, nil
		}

		return os.Stderr, nil, nil
	case File:
		if len(sink.Args) == 0 {
			return nil, nil, ErrArgForFileNotProvided
		}

		var path string

		switch val := sink.Args[0].(type) {
		case string:
			path = val
		case fmt.Stringer:
			path = val.String()
		default:
			return nil, nil, ErrUnexpectedArgForFileType
		}

		f, err := utils.CreateLogFile(path)
		if err != nil {
			return nil, nil, err
		}

		return f, func() error {
			return f.Close()
		}, nil
	default:
		return nil, nil, ErrInvalidSinkType
	}
}
