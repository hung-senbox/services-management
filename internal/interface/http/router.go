package http

import (
	"services-management/internal/interface/http/handler"
	"services-management/internal/interface/http/route"
	"services-management/internal/interface/middleware"
	"services-management/pkg/gateway"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRouter sets up the Fiber router
func SetupRouter(
	serviceGroupHandler *handler.ServiceGroupHandler,
	serviceHandler *handler.ServiceHandler,
	auditMiddleware *middleware.AuditMiddleware,
	userGateway gateway.UserGateway,
) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Services Management v1.0",
	})

	// Apply global middlewares
	app.Use(fiberLogger.New())
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
	route.SetUpServiceGroupRoutes(app, serviceGroupHandler, userGateway)

	// Services routes
	route.SetUpServiceRoutes(app, serviceHandler)

	return app
}
