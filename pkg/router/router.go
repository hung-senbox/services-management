package router

import (
	"services-management/internal/sv_management/handler"
	"services-management/internal/sv_management/repository"
	"services-management/internal/sv_management/route"
	service "services-management/internal/sv_management/services"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(consulClient *api.Client, serviceCollection *mongo.Collection, serviceGroupCollection *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// gateway
	//userGateway := gateway.NewUserGateway("go-main-service", consulClient)

	// services group
	serviceGroupRepo := repository.NewServiceGroupRepository(serviceGroupCollection)
	serviceGroupService := service.NewSVGroupService(serviceGroupRepo)
	serviceGroupHandler := handler.NewServiceGroupHandler(serviceGroupService)

	// services
	serviceRepo := repository.NewServiceRepository(serviceCollection)
	svManagementService := service.NewSvManagementService(serviceRepo, serviceGroupRepo)
	serviceHandler := handler.NewServiceHandler(svManagementService)

	// Register routes
	route.RegisterServiceRoutes(r, serviceHandler, serviceGroupHandler)
	//route.RegisterRegionRoutes(r, regionHandler)
	return r
}
