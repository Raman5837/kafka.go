package routes

import (
	"github.com/Raman5837/kafka.go/app/handler"
	"github.com/gofiber/fiber/v2"
)

// All Routes Related To Producer
func ProducerRoutes(app *fiber.App) {

	route := app.Group("/api/v1/produce")
	route.Post("/message", handler.ProduceMessageHandler)

}
