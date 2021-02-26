package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"shortify/handler"
	"shortify/middleware"
)

func Register(app *fiber.App) {
	// Middleware
	pub := app.Group("", logger.New())
	api := app.Group("/api", logger.New())

	// Public
	pub.Get("/health", handler.HealthCheck)
	pub.Get("/s/:token", handler.GetOriginalURL)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/token", handler.Login)

	// URLs
	url := api.Group("/urls", middleware.AuthRequired())
	url.Get("", handler.GetAllURLs)
	url.Post("", handler.CreateURL)
	url.Get("/:uuid", handler.GetURL)
	url.Delete("/:uuid", handler.DeleteURL)
}