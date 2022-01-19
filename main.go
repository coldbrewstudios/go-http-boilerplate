package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fibreLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"go-http-boilerplate/app/routing"
	"go-http-boilerplate/pkg/logging"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	logger := logging.CreateLogger("service")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		logger.Error("Error while reading config file: ", err)
	}

	app := fiber.New()
	setupMiddleware(app)
	setupRouting(app)
	StartServer(app, logger)
}

// setupMiddleware
// Sets up all middleware for Fiber
func setupMiddleware(app *fiber.App) {
	var frontendAssets fs.FS

	env, ok := viper.Get("ENVIRONMENT").(string)

	if !ok {
		log.Fatalf("Could not find port ENV variable")
	}

	if env == "development" {
		frontendAssets = getDevFrontendAssets()
	} else {
		frontendAssets = getProdFrontendAssets()
	}

	app.Use(fibreLogger.New(fiberLogger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(filesystem.New(filesystem.Config{
		Root: http.FS(frontendAssets),
	}))
}

// setupMiddleware
// Sets up all routes and base groups for the application
func setupRouting(app *fiber.App) {
	baseRouter := app.Group("/api/v1")

	routing.CreatePublicRoutes(baseRouter)
}

// getProdFrontendAssets
// Uses frontend to create an embed file system for build time to bundle assets together
// in production
func getProdFrontendAssets() fs.FS {
	f, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		panic(err)
	}

	return f
}

// getProdFrontendAssets
// Serves the fronted/dist folder as the underlying file system. If this is being changed it will also pick
// these changes up
func getDevFrontendAssets() fs.FS {
	return os.DirFS("frontend/dist")
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(app *fiber.App, logger logging.AppLogger) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	StartServer(app, logger)

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(app *fiber.App, logger logging.AppLogger) {
	port, ok := viper.Get("PORT").(string)

	if !ok {
		log.Fatalf("Could not find port ENV variable")
	}

	logger.Info("Application Bootstrap Complete")

	listenErr := app.Listen(":" + port)

	if listenErr != nil {
		logger.Error("Error during listen", listenErr)
	}
}
