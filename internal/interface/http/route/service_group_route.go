package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senbox/services-management/internal/interface/http/handler"
)

func SetUpServiceGroupRoutes(app *fiber.App, serviceGroupHandler *handler.ServiceGroupHandler) {
	// API v1 routes
	api := app.Group("/api/v1")

	// admin routes
	admin := api.Group("/admin")

	// Service Groups routes
	serviceGroups := admin.Group("/service-groups")
	serviceGroups.Post("/", serviceGroupHandler.CreateServiceGroup)
	serviceGroups.Get("/", serviceGroupHandler.GetAllServiceGroups)
	serviceGroups.Get("/:id", serviceGroupHandler.GetServiceGroupByID)
	serviceGroups.Put("/:id", serviceGroupHandler.UpdateServiceGroup)
	serviceGroups.Delete("/:id", serviceGroupHandler.DeleteServiceGroup)
}
