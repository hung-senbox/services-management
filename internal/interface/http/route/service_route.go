package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senbox/services-management/internal/interface/http/handler"
)

func SetUpServiceRoutes(app *fiber.App, serviceHandler *handler.ServiceHandler) {
	// API v1 routes
	api := app.Group("/api/v1")

	// admin routes
	admin := api.Group("/admin")

	// Services routes
	services := admin.Group("/services")
	services.Post("/", serviceHandler.CreateService)
	services.Get("/", serviceHandler.GetAllServices)
	services.Get("/:id", serviceHandler.GetServiceByID)
	services.Get("/group/:groupId", serviceHandler.GetServicesByGroupID)
	services.Put("/:id", serviceHandler.UpdateService)
	services.Delete("/:id", serviceHandler.DeleteService)
}
