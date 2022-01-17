package main

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	fibreLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fibreRecover "github.com/gofiber/fiber/v2/middleware/recover"
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

	logger := logging.CreateLogger(env)

	port, ok := viper.Get("PORT").(string)

	if !ok {
		log.Fatalf("Could not find port ENV variable")
	}

	app := fiber.New()
	httpLogFile := setupHttpLogs()
	loggerConfig := fibreLogger.Config{Output: httpLogFile}
	setupMiddleware(app, loggerConfig, env)
	setupRouting(app)

	listenErr := app.Listen(":" + port)

	if listenErr != nil {
		logger.Error("Error during listen", listenErr)
	}
}

func setupRouting(app *fiber.App) {
	baseRouter := app.Group("/api/v1")
	routing.CreatePublicRoutes(baseRouter)
}

func setupMiddleware(app *fiber.App, lc fibreLogger.Config, env string) {
	var frontendAssets fs.FS

	if env == "development" {
		frontendAssets = getDevFrontendAssets()
	} else {
		frontendAssets = getProdFrontendAssets()
	}

	app.Use(fibreRecover.New())
	app.Use(fibreLogger.New(lc))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(filesystem.New(filesystem.Config{
		Root: http.FS(frontendAssets),
	}))
}

func setupHttpLogs() *os.File {
	file, err := os.OpenFile("./var/log/http.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return file
}

func getProdFrontendAssets() fs.FS {
	f, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		panic(err)
	}

	return f
}

func getDevFrontendAssets() fs.FS {
	return os.DirFS("frontend/dist")
}
