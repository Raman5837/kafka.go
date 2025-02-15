package settings

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Raman5837/kafka.go/base/database"
	"github.com/Raman5837/kafka.go/base/middleware"
	"github.com/Raman5837/kafka.go/base/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Returns A New Instance Of Fiber Application
func NewFiberApp() *fiber.App {

	config := fiber.Config{AppName: "Kafka.go", Prefork: false, ServerHeader: "Kafka.go"}

	allowedMethods := strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodPatch,
		fiber.MethodDelete,
		fiber.MethodOptions,
	}, ",")

	allowedHeaders := "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin"

	app := fiber.New(config)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     allowedHeaders,
		AllowMethods:     allowedMethods,
		AllowOrigins:     "http://localhost.com",
	}))

	// Attaching All Middlewares
	AddMiddleware(app)

	// Connect All Databases
	ConnectDataBase()

	return app
}

// Add All The Defined Middleware (Ordering Is Important)
func AddMiddleware(app *fiber.App) {

	// 1. Add A Request Id To All The Requests
	app.Use(middleware.RequestId())

	// 2. Logger All The Requests
	app.Use(middleware.APILogger())

	// 3. API Error Handler
	app.Use(middleware.ErrorHandler())
}

// Connect To DataBase
func ConnectDataBase() {

	// Initialize Database Connections
	if dbError := database.EstablishConnection(); dbError != nil {
		utils.Logger.Fatal(dbError, "Error While Connecting To Database: ")
	}

}

/*
Gracefully Shutdown The Application

The server will wait for all the active connections to process, and will not accept new connections.
*/
func GracefulShutdownHandler(app *fiber.App, shutdown chan os.Signal) {

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		utils.Logger.Info("Shutting Down Gracefully...")

		// Close any other resources or perform cleanup before shutting down the app
		if exception := app.Shutdown(); exception != nil {
			utils.Logger.Error(exception, "Error During Shutdown")
		}

		// Close the shutdown channel to signal that the shutdown process is complete
		close(shutdown)
	}()

}

/*
Performs Cleanup Tasks Before Shutdown And After Program Exits.

It Closes All The DB Connections And Redis Connections As Of Now.
*/
func InitiateCleanupProcess() {

	sqlite := database.DBManager.SqliteDB

	if sqlite != nil {

		if db, err := sqlite.DB(); err == nil {

			if closingError := db.Close(); closingError != nil {
				utils.Logger.Error(closingError, "Error While Closing Sqlite DB Connection")
			} else {
				utils.Logger.Info("Sqlite DB Connection Closed")
			}
		}

	}
}
