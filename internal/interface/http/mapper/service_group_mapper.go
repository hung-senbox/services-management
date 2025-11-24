package mapper

import (
	"github.com/senbox/services-management/internal/domain/entity"
	"github.com/senbox/services-management/internal/interface/http/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToServiceGroupEntity(req *dto.CreateServiceGroupRequest) *entity.ServiceGroup {
	return &entity.ServiceGroup{
		Name:        req.Name,
		Order:       req.Order,
		IsActive:    req.IsActive,
		Description: req.Description,
		Url:         req.Url,
		Icon:        req.Icon,
	}
}

func ToServiceGroupResponse(sg *entity.ServiceGroup) *dto.ServiceGroupResponse {
	return &dto.ServiceGroupResponse{
		ID:          sg.ID.Hex(),
		Name:        sg.Name,
		Order:       sg.Order,
		IsActive:    sg.IsActive,
		Description: sg.Description,
		Icon:        sg.Icon,
		Url:         sg.Url,
		CreatedAt:   sg.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   sg.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func ToServiceGroupEntityFromUpdate(id string, req *dto.UpdateServiceGroupRequest) (*entity.ServiceGroup, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return &entity.ServiceGroup{
		ID:          objectID,
		Name:        req.Name,
		Order:       req.Order,
		IsActive:    req.IsActive,
		Description: req.Description,
		Url:         req.Url,
		Icon:        req.Icon,
	}, nil
}
