package routes

import (
	"github.com/Raman5837/kafka.go/app/handler"
	"github.com/gofiber/fiber/v2"
)

// All Routes Related To ConsumerGroup
func ConsumerGroupRoutes(app *fiber.App) {

	route := app.Group("/api/v1/consumer-group")
	route.Post("/add", handler.AddNewConsumerHandler)
}
