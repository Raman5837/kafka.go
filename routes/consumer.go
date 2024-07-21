package routes

import (
	"github.com/Raman5837/kafka.go/app/handler"
	"github.com/gofiber/fiber/v2"
)

// All Routes Related To Consumer
func ConsumerRoutes(app *fiber.App) {

	route := app.Group("/api/v1/consumer")
	route.Post("/add", handler.AddNewConsumerHandler)
	route.Get("/consume", handler.ConsumeMessageHandler)
}
