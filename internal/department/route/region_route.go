package route

import (
	"department-service/internal/department/handler"
	"department-service/internal/middleware"

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
