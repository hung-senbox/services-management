package route

import (
	"services-management/internal/department/handler"
	"services-management/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRegionRoutes(r *gin.Engine, h *handler.RegionHandler) {
	adminGroup := r.Group("/api/v1/admin")
	adminGroup.Use(middleware.Secured())
	{
		departmentsAdmin := adminGroup.Group("/regions")
		{
			departmentsAdmin.POST("", h.CreateRegion)
			departmentsAdmin.PUT("/:region_id", h.UpdateRegionName)
		}
	}
}
