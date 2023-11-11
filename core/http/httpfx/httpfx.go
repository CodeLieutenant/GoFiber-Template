package httpfx

import (
	"context"
	"fmt"
	"net/http"
	"time"

	corehttp "github.com/BrosSquad/GoFiber-Boilerplate/core/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

func Module(addr, appName string, displayInfo bool, routes ...fx.Option) fx.Option {
	return fx.Module("fiber-app",
		append(routes,
			fx.Provide(func(lc fx.Lifecycle) *fiber.App {
				app := corehttp.CreateApplication(appName, displayInfo)
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							_ = app.Listen(addr)
						}()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						newCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
						defer cancel()
						return app.ShutdownWithContext(newCtx)
					},
				})

				return app
			}),
		)...,
	)
}

func Get(path string, handlers ...any) func(string) fx.Option {
	return Route(http.MethodGet, path, handlers...)
}

func Post(path string, handlers ...any) func(string) fx.Option {
	return Route(http.MethodPost, path, handlers...)
}

func Put(path string, handlers ...any) func(string) fx.Option {
	return Route(http.MethodPut, path, handlers...)
}

func Patch(path string, handlers ...any) func(string) fx.Option {
	return Route(http.MethodPatch, path, handlers...)
}

func Delete(path string, handlers ...any) func(string) fx.Option {
	return Route(http.MethodDelete, path, handlers...)
}

func Route(method, path string, handlers ...any) func(string) fx.Option {
	return func(prefix string) fx.Option {
		return fx.Module(fmt.Sprintf("route-%s-%s%s", method, prefix, path), append(
			lo.Map(handlers, func(a any, _ int) fx.Option {
				return fx.Provide(fx.Annotate(
					a,
					fx.ResultTags(fmt.Sprintf(`group:"handlers-%s-%s%s"`, method, prefix, path)),
				))
			}),
			fx.Invoke(
				fx.Annotate(
					func(router fiber.Router, handlers []fiber.Handler) {
						router.Add(method, path, handlers...)
					},
					fx.ParamTags(
						fmt.Sprintf(`name:"router-%s"`, prefix),
						fmt.Sprintf(`group:"handlers-%s-%s%s"`, method, prefix, path),
					),
				),
			),
		)...)
	}
}

func Routes(prefix string, routes ...func(string) fx.Option) fx.Option {
	return fx.Module(
		fmt.Sprintf("routes-%s", prefix),
		append(
			lo.Map(routes, func(fn func(string) fx.Option, _ int) fx.Option {
				return fn(prefix)
			}),
			fx.Provide(
				fx.Annotate(
					func(app *fiber.App) fiber.Router {
						return app.Group(prefix)
					},
					fx.ResultTags(fmt.Sprintf(`name:"router-%s"`, prefix)),
				),
			),
		)...,
	)
}
