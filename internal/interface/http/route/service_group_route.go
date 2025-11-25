package route

import (
	"services-management/internal/interface/http/handler"
	"services-management/internal/interface/middleware"
	"services-management/pkg/gateway"

	"github.com/gofiber/fiber/v2"
)

func SetUpServiceGroupRoutes(app *fiber.App, serviceGroupHandler *handler.ServiceGroupHandler, userGateway gateway.UserGateway) {
	// API v1 routes
	api := app.Group("/api/v1")
	api.Use(middleware.Secured(userGateway))
	// admin routes
	admin := api.Group("/admin")

	// Service Groups routes
	serviceGroups := admin.Group("/service-groups")
	serviceGroups.Post("", serviceGroupHandler.CreateServiceGroup)
	serviceGroups.Get("", serviceGroupHandler.GetAllServiceGroups)
	serviceGroups.Post("/migrate", serviceGroupHandler.MigrateServiceGroup)
	serviceGroups.Get("/:id", serviceGroupHandler.GetServiceGroupByID)
	serviceGroups.Put("/:id", serviceGroupHandler.UpdateServiceGroup)
	serviceGroups.Delete("/:id", serviceGroupHandler.DeleteServiceGroup)
}
