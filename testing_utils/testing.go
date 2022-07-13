package testing_utils

import (
	"errors"
	"os"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/config"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/container"
	httpapp "github.com/BrosSquad/GoFiber-Boilerplate/pkg/http"
	"github.com/BrosSquad/GoFiber-Boilerplate/pkg/utils"
)

func GetValidator() (*validator.Validate, ut.Translator) {
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	englishTranslations, _ := uni.GetTranslator("en")

	return v, englishTranslations
}

func CreateApplication() (*fiber.App, *container.Container) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath, err := findConfig(wd)
	if err != nil {
		panic(err)
	}

	cfg := viper.New()

	cfg.SetConfigName("config")
	cfg.AddConfigPath(configPath)
	cfg.SetConfigType("yaml")

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

	c := container.New(cfg, false, "info", config.Testing)
	return httpapp.CreateApplication(c, constants.AppName, config.Testing, false, false, nil), c
}

func findConfig(workingDir string) (string, error) {
	for entries, err := os.ReadDir(workingDir); err == nil; {
		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == "config" {
				return workingDir, nil
			}
		}

		workingDir, err = utils.GetAbsolutePath(workingDir + "/..")

		if err != nil {
			return "", err
		}

		entries, err = os.ReadDir(workingDir)
	}

	return "", errors.New("config file not found")
}
