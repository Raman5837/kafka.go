package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// API Error Handler Middleware
func ErrorHandler() fiber.Handler {

	return recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	)
}

func DefaultErrorHandler(context *fiber.Ctx, apiError error) error {

	// Default Status Code
	statusCode := fiber.StatusInternalServerError

	// Retrieve The Custom Status Code If It's A *fiber.Error
	var newErrorInstance *fiber.Error
	if errors.As(apiError, &newErrorInstance) {
		statusCode = newErrorInstance.Code
	}

	return context.Status(statusCode).SendString("Something Went Wrong")

}
