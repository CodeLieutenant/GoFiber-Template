package testing_utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/pkg/config"
	"github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/pkg/constants"
	httpapp "github.com/BrosSquad/GoFiber-GoFiber-Boilerplate/pkg/http"
)

func GetValidator() (*validator.Validate, ut.Translator) {
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	englishTranslations, _ := uni.GetTranslator("en")

	return v, englishTranslations
}


func CreateApplication() *fiber.App {
	return httpapp.CreateApplication(constants.AppName, config.Testing, false, nil)
}
