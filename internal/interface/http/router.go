package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senbox/services-management/internal/interface/http/handler"
	"github.com/senbox/services-management/internal/interface/http/route"
	"github.com/senbox/services-management/internal/interface/middleware"
)

// SetupRouter sets up the Fiber router
func SetupRouter(
	serviceGroupHandler *handler.ServiceGroupHandler,
	serviceHandler *handler.ServiceHandler,
	auditMiddleware *middleware.AuditMiddleware,
) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Services Management v1.0",
	})

	// Apply global middlewares
	app.Use(middleware.LoggingMiddleware())
	app.Use(middleware.CORSMiddleware())
	app.Use(auditMiddleware.Log())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Service Groups routes
	route.SetUpServiceGroupRoutes(app, serviceGroupHandler)

	// Services routes
	route.SetUpServiceRoutes(app, serviceHandler)

	return app
}
