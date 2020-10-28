package main

import (
	"github.com/spf13/viper"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Setting up Viper
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/gohack")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	app := fiber.New(fiber.Config{
		Prefork: viper.GetBool("http.prefork"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(viper.GetString("http.address")); err != nil {
		log.Fatalf("Error while starting application: %v", err)
	}
}
