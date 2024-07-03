package main

import (
	"os"

	"github.com/Raman5837/kafka.go/base/settings"
	"github.com/Raman5837/kafka.go/base/utils"
	"github.com/Raman5837/kafka.go/routes"
)

// Initializing All The Global Variables
func init() {
	utils.InitLogger()
}

// Entrypoint Of The Application
func main() {

	// New Fiber App.
	app := settings.NewFiberApp()

	// Registering All Routes
	routes.RegisterAllRoutes(app)

	// Graceful Shutdown
	shutdown := make(chan os.Signal, 1)
	settings.GracefulShutdownHandler(app, shutdown)

	// This Will Get Execute, After The Main Function
	defer settings.InitiateCleanupProcess()

	// Listening on PORT Defined In ENV
	serverPort := ":" + utils.GetEnv("SERVER_PORT")
	if serverError := app.Listen(serverPort); serverError != nil {
		utils.Logger.Fatal(serverError, "Error Starting Go-Kafka App")
	}

}
