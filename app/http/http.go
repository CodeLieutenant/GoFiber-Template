package http

import (
	"context"

	"github.com/BrosSquad/GoFiber-Boilerplate/app/constants"
	"github.com/BrosSquad/GoFiber-Boilerplate/app/container"
	corehttp "github.com/BrosSquad/GoFiber-Boilerplate/core/http"
	"github.com/gofiber/fiber/v2"
)

func CreateApplication(ctx context.Context, c *container.Container, displayInfo bool) *fiber.App {
	return corehttp.CreateApplication(ctx, constants.AppName, c, displayInfo, routes)
}
