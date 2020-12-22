package middleware

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Timeout(d time.Duration, h ...fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), d)
		ch := make(chan error, 1)

		defer cancel()
		defer close(ch)

		go func() {
			if len(h) > 0 {
				ch <- h[0](c)
			} else {
				ch <- c.Next()
			}
		}()

		select {
		case err := <-ch:
			return err
		case <-ctx.Done():
			return fiber.ErrRequestTimeout
		}
	}
}
