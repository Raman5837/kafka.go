package routes

import "github.com/gofiber/fiber/v2"

// Register All Routes From All Sub Modules(App)
func RegisterAllRoutes(app *fiber.App) {

	ServiceSanityRoutes(app)
	ConsumerGroupRoutes(app)
	ConsumerRoutes(app)
	TopicRoutes(app)

}
