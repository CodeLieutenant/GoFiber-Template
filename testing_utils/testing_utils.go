package testing_utils

import (
	"context"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"github.com/nano-interactive/go-utils/environment"
	nanoTesting "github.com/nano-interactive/go-utils/testing"

	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/config"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/constants"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/container"
	"github.com/nano-interactive/GoFiber-Boilerplate/pkg/http/handlers"

	httpapp "github.com/nano-interactive/GoFiber-Boilerplate/pkg/http"
)

type Modifier func(*config.Config)

func CreateApplication(t *testing.T, modifier ...Modifier) (*fiber.App, *container.Container) {
	t.Helper()

	return nanoTesting.CreateApplicationFunc(func(ctx context.Context, v *viper.Viper) (*fiber.App, *container.Container) {
		cfg, err := config.NewWithViper(v)
		if err != nil {
			t.Fatal(err)
			t.FailNow()
		}

		for _, mod := range modifier {
			mod(cfg)
		}

		c := container.New(context.Background(), cfg, true, "debug", environment.Testing)
		return httpapp.CreateApplication(c, constants.AppName, environment.Testing, false, false,
			handlers.Error(zerolog.Nop(), c.GetTranslator())), c
	})
}
