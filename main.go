package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/Gridmax/Fader/utility/configload"
)

func main() {
    app := fiber.New()
    config, err := configload.LoadConfig("config.yaml")
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })

    app.Listen(":"+config.ServicePort)
}
