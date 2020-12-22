package main

import (
	"context"
	"flag"
	"io"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/api"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/api/routes"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/container"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/database"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/handlers"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/logging"
)

func createLogFile(path string) (file io.WriteCloser, err error) {
	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, 0744); err != nil {
		return nil, err
	}

	file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

	if err != nil {
		return nil, err
	}

	return
}

func main() {
	loggingLevel := flag.String("logging", "debug", "Global logging level")
	flag.Parse()

	zerolog.SetGlobalLevel(logging.Parse(*loggingLevel))
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/nanorequestdecomposer")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	log.Debug().Msg("Loading configuration files\n")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Fatal error config file: %s \n", err)
	}

	logFile, err := createLogFile(viper.GetString("logging.file"))

	if err != nil {
		log.Fatal().Err(err).Msg("Error while opening log file")
	}

	dbLogFile, err := createLogFile(viper.GetString("database.logfile"))

	if err != nil {
		log.Fatal().Err(err).Msg("Error while opening database log file")
	}

	logger := logging.New(
		viper.GetString("logging.level"),
		viper.GetBool("logging.console"),
		logFile,
	)

	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator(viper.GetString("locale"))
	validate := validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		logger.Fatal().Err(err).Msg("Error while registering english translations")
	}

	db, err := database.ConnectDB(database.Config{}, dbLogFile, viper.GetBool("debug"))

	if err != nil {
		log.Fatal().Err(err).Msg("Error while connecting to database")
	}

	diContainer := container.Container{
		Ctx:       ctx,
		Logger:    &logger,
		Validator: validate,
		DB:        db,
	}

	go func(cancel *context.CancelFunc) {
		s := <-signalCh
		logger.Info().Msgf("Shutting down... Signal: %s\n", s)
		(*cancel)()
	}(&cancel)

	logger.Debug().Msg("Starting HTTP Api")

	provider := api.NewFiberAPI(
		viper.GetString("http.address"),
		viper.GetBool("http.prefork"),
		viper.GetBool("debug"),
		handlers.Error(diContainer.Logger, trans),
		routes.RegisterRouter,
	)

	go func() {
		<-ctx.Done()

		if err := provider.Close(); err != nil {
			logger.Error().
				Err(err).
				Msg("Error while shutting down application\n")
		}

		if err := database.Close(); err != nil {
			diContainer.Logger.Err(err).Msg("Error while closing sql database file")
		}

		if err := logFile.Close(); err != nil {
			diContainer.Logger.Err(err).Msg("Error while closing log file")
		}

		if err := dbLogFile.Close(); err != nil {
			diContainer.Logger.Err(err).Msg("Error while closing database log file")
		}

	}()

	if err := provider.Register(&diContainer); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Error while starting the app")
	}
}
