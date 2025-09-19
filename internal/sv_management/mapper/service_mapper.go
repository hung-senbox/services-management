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
