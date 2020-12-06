package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/api"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/handlers"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/logging"
)

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

	logger := logging.New(viper.GetString("logging.level"))

	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator(viper.GetString("locale"))
	validate := validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		logger.Fatal().Err(err).Msg("Error while registering english translations")
	}

	container := api.Container{
		Ctx:       ctx,
		Logger:    logger,
		Validator: validate,
	}

	go func(cancel *context.CancelFunc) {
		s := <-signalCh
		logger.Info().Msgf("Shutting down... Signal: %s\n", s)
		(*cancel)()
	}(&cancel)

	logger.Debug().Msg("Starting HTTP Api")

	provider := api.NewFiberAPI(
		viper.GetBool("http.prefork"),
		viper.GetBool("debug"),
		handlers.Error(trans),
	)

	go func() {
		<-ctx.Done()

		if err := provider.Close(); err != nil {
			logger.Error().
				Err(err).
				Msg("Error while shutting down application\n")
		}
	}()

	if err := provider.Register(&container); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Error while starting the app")
	}
}
