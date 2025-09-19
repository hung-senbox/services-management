package route

import (
	"services-management/internal/middleware"
	"services-management/internal/sv_management/handler"

	"github.com/gin-gonic/gin"
)

func RegisterServiceRoutes(r *gin.Engine, sh *handler.ServiceHandler, sgh *handler.ServiceGroupHandler) {
	// Admin routes
	adminGroup := r.Group("/api/v1/admin", middleware.Secured(), middleware.RequireAdmin())

	// Service routes
	services := adminGroup.Group("/services")
	{
		services.POST("", sh.Upload)

		// Service group routes
		groups := services.Group("/groups")
		{
			groups.POST("", sgh.Upload)
		}
	}
}
