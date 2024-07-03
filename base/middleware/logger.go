package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// API Logger Handler
func APILogger() fiber.Handler {

	loggingFormat := `{"pid": ${pid}, "requestId": "${locals:requestid}", "time": "${time}", url: ${host}${url}, "status": ${status}, "latency": "${latency}", "method": "${method}", "path": "${path}", "body": "${body}", "queryParams": "${queryParams}"}` + "\n"

	return logger.New(
		logger.Config{
			Format:     loggingFormat,
			TimeZone:   "Asia/Kolkata",
			TimeFormat: "2006-01-02T15:04:05-0700",
		},
	)
}
