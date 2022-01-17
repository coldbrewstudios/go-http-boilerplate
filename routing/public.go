package routing

import (
	"github.com/gofiber/fiber/v2"
	"go-http-boilerplate/handlers"
)

func CreatePublicRoutes(rg fiber.Router) {
	rg.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("OK")
	})

	rg.Get("/law", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(handlers.GetRandomLaw())
	})
}