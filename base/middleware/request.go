package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// Add Request Id In All Requests
func RequestId() fiber.Handler {

	return requestid.New()
}
