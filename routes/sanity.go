package routes

import (
	"github.com/Raman5837/kafka.go/base/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Handler For The Base Endpoint
func BaseHandler(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(utils.HttpResponseOK(nil, "Kafka.go Is Ready To Serve!"))
}

// Health Check API Handler
func HealthCheck(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(utils.HttpResponseOK(nil, "Kafka.go Is Up And Running"))
}

// Handler To Monitor Service Resources
func ResourceMonitor() fiber.Handler {
	return monitor.New(monitor.Config{Title: "Kafka.go Metrics Page"})
}

// Basic Routes (Currently Contains health And Metrics View)
func ServiceSanityRoutes(app *fiber.App) {

	route := app.Group("/")

	// Routes For GET Methods
	route.Get("/", BaseHandler)
	route.Get("/api/check", HealthCheck)
	route.Get("/api/monitor", ResourceMonitor())
}
