package testutils

import (
	"context"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	nanotesting "github.com/nano-interactive/go-utils/testing"
	nanofibertesting "github.com/nano-interactive/go-utils/testing/fiber"
	"github.com/samber/lo"
	"github.com/spf13/viper"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/config"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/container"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/http"
)

func GetConfig(tb testing.TB) config.Config {
	tb.Helper()
	return nanotesting.GetConfig[config.Config](tb, func(v *viper.Viper) (config.Config, error) {
		cfg, err := config.NewWithViper(v)
		if err != nil {
			return config.Config{}, err
		}

		cfg.Logging.PrettyPrint = true

		return cfg, nil
	})
}

func AppTester[T any](tb testing.TB, items ...any) *nanofibertesting.GoFiberSender[T] {
	tb.Helper()
	sender, _, _ := App[T](tb, false, items...)
	return sender
}

func App[T any](tb testing.TB, followRedirects bool, items ...any) (*nanofibertesting.GoFiberSender[T], *container.Container, *fiber.App) {
	tb.Helper()

	creator := nanotesting.AppCreaterFunc[*fiber.App, *container.Container](func(ctx context.Context, v *viper.Viper) (*fiber.App, *container.Container) {
		conf, found := lo.Find(items, func(l any) bool {
			_, ok := l.(config.Config)

			return ok
		})

		cfg, err := config.NewWithViper(v)
		if err != nil {
			tb.Errorf("Failed to parse config file: %v", v)
			tb.FailNow()
		}

		if found {
			cfg = conf.(config.Config)
		}
		di := container.New(ctx, cfg)
		app := http.CreateApplication(ctx, di, false)

		tb.Cleanup(func() {
			if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
				tb.Errorf("Failed to close APP Server: %v", err)
			}

			if err := di.Close(); err != nil {
				tb.Errorf("Failed to close DI Container: %v", err)
			}
		})

		return nil, nil
	})

	app, di := nanotesting.CreateApplication[*fiber.App, *container.Container](tb, creator)
	sender := nanofibertesting.New[T](tb, app, followRedirects)

	return sender, di, app
}
