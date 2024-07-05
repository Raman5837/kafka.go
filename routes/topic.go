package routes

import (
	"github.com/Raman5837/kafka.go/app/handler"
	"github.com/gofiber/fiber/v2"
)

// All Routes Related To Topic
func TopicRoutes(app *fiber.App) {

	route := app.Group("/api/v1/topic")
	route.Get("/", handler.GetTopicHandler)
	route.Post("/create", handler.CreateTopicHandler)

}
