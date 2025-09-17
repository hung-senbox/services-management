package router

import (
	"services-management/internal/department/handler"
	"services-management/internal/department/repository"
	"services-management/internal/department/route"
	"services-management/internal/department/service"
	"services-management/internal/gateway"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(consulClient *api.Client, departmentCollection *mongo.Collection, regionCollection *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// gateway
	userGateway := gateway.NewUserGateway("go-main-service", consulClient)
	menuGateway := gateway.NewMenuGateway("go-main-service", consulClient)

	// region
	regionRepo := repository.NewRegionRepository(regionCollection)
	regionService := service.NewRegionService(regionRepo, userGateway)
	regionHandler := handler.NewRegionHandler(regionService)

	// department
	departmentRepo := repository.NewDepartmentRepository(departmentCollection)
	departmentService := service.NewDepartmentService(departmentRepo, userGateway, menuGateway, regionRepo)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	// Register routes
	route.RegisterDepartmentRoutes(r, departmentHandler, regionHandler)
	//route.RegisterRegionRoutes(r, regionHandler)
	return r
}
