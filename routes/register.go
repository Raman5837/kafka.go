package routes

import "github.com/gofiber/fiber/v2"

// Register All Routes From All Sub Modules(App)
func RegisterAllRoutes(app *fiber.App) {

	ServiceSanityRoutes(app)
	ConsumerRoutes(app)
	TopicRoutes(app)

}
