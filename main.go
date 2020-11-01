package main

import (
	"context"
	"flag"
	"gohack/api"
	"gohack/database"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func parseLogginLevel(level string) zerolog.Level {
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

func createLogger(writers ...io.Writer) zerolog.Logger {
	if viper.GetBool("logging.console") {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	logger := zerolog.New(io.MultiWriter(writers...)).With().Logger()
	logger.Level(parseLogginLevel(viper.GetString("logging.level")))

	return logger
}

func main() {
	logginLevel := flag.String("logging", "debug", "Global logging level")
	flag.Parse()
	zerolog.SetGlobalLevel(parseLogginLevel(*logginLevel))
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().Msg("Starting application...\n")
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)

	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/gohack")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	log.Debug().Msg("Loading configuration files\n")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Fatal error config file: %s \n", err)
	}

	log.Debug().Msg("Connecting to database\n")
	db, err := database.ConnectDB(database.Config{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.dbname"),
		Port:     uint16(viper.GetUint32("database.port")),
		TimeZone: viper.GetString("database.timezone"),
		SslMode:  viper.GetBool("database.sslmode"),
	})

	container := api.Container{
		DB:     db,
		Ctx:    ctx,
		Logger: createLogger(), // TODO: Add files for logging
	}

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Fatal error database file\n")
	}

	go func(ctx context.Context) {
		log.Debug().Msg("Starting HTTP Api")
		provider := api.NewFiberApi(viper.GetBool("http.prefork"))
		provider.Register(&container)
		<-ctx.Done()
		done <- provider.Close()
	}(ctx)

	s := <-signalCh
	cancel()
	err = <-done

	if err != nil {
		log.Error().
			Err(err).
			Msg("Error while shutting down application\n")
	}

	log.Info().Msgf("Shutting down... Signal: %s\n", s)
}
