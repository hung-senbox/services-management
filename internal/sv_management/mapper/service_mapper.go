package mapper

import (
	"services-management/internal/sv_management/dto/response"
	"services-management/internal/sv_management/model"
)

func MapServiceToServiceResDto(service model.Service) *response.ServiceResDto {
	return &response.ServiceResDto{
		ID:    service.ID.Hex(),
		Title: service.Title,
		Order: service.Order,
		Url:   service.Url,
	}
}

func MapServicesResponse(groups []*model.ServiceGroup, services []*model.Service) []*response.ServicesResponse {
	// Gom service theo group
	serviceMap := make(map[string][]response.ServiceResDto)
	for _, svc := range services {
		serviceMap[svc.GroupID] = append(serviceMap[svc.GroupID], response.ServiceResDto{
			ID:    svc.ID.Hex(),
			Title: svc.Title,
			Url:   svc.Url,
			Order: svc.Order,
		})
	}

	// Build response
	var result []*response.ServicesResponse
	for _, g := range groups {
		res := &response.ServicesResponse{
			Group: response.ServiceGroupResponse{
				ID:    g.ID.Hex(),
				Title: g.Title,
				Order: g.Order,
			},
			Services: serviceMap[g.ID.Hex()],
		}
		result = append(result, res)
	}

	return result
}
