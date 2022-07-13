package container

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog/log"
)

func (c *Container) GetValidator() *validator.Validate {
	if c.validator == nil {
		c.validator = validator.New()

		if err := entranslations.RegisterDefaultTranslations(c.validator, c.GetTranslator()); err != nil {
			log.
				Fatal().
				Err(err).
				Msg("Error while registering english translations")
		}
	}

	return c.validator
}

func (c *Container) GetTranslator() ut.Translator {
	if c.translator == nil {
		english := en.New()
		uni := ut.New(english, english)

		translator, found := uni.GetTranslator("en")

		if !found {
			log.
				Fatal().
				Msg("Locale is not found: en")
		}

		c.translator = translator
	}

	return c.translator
}
