package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fibreLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"go-http-boilerplate/pkg/logging"
	"go-http-boilerplate/routing"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	env, ok := viper.Get("ENVIRONMENT").(string)

	if !ok {
		log.Fatalf("Could not find port ENV variable")
	}

	logger := logging.CreateLogger(env, "service")
	port, ok := viper.Get("PORT").(string)

	if !ok {
		log.Fatalf("Could not find port ENV variable")
	}

	app := fiber.New()
	setupMiddleware(app, env)
	setupRouting(app)
	logger.Info("Application Bootstrap Complete")

	listenErr := app.Listen(":" + port)

	if listenErr != nil {
		logger.Error("Error during listen", listenErr)
	}
}

// setupMiddleware
// Sets up all middleware for Fiber
func setupMiddleware(app *fiber.App, env string) {
	var frontendAssets fs.FS

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
