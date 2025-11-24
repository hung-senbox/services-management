package mapper

import (
	"github.com/senbox/services-management/internal/domain/entity"
	"github.com/senbox/services-management/internal/interface/http/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToServiceEntity(req *dto.CreateServiceRequest) (*entity.Service, error) {
	serviceGroupID, err := primitive.ObjectIDFromHex(req.ServiceGroupID)
	if err != nil {
		return nil, err
	}

	return &entity.Service{
		ServiceGroupID: serviceGroupID,
		Name:           req.Name,
		IsActive:       req.IsActive,
		Description:    req.Description,
	}, nil
}

func ToServiceResponse(s *entity.Service) *dto.ServiceResponse {
	return &dto.ServiceResponse{
		ID:             s.ID.Hex(),
		ServiceGroupID: s.ServiceGroupID.Hex(),
		Name:           s.Name,
		IsActive:       s.IsActive,
		Description:    s.Description,
		CreatedAt:      s.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      s.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func ToServiceEntityFromUpdate(id string, req *dto.UpdateServiceRequest) (*entity.Service, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	serviceGroupID, err := primitive.ObjectIDFromHex(req.ServiceGroupID)
	if err != nil {
		return nil, err
	}

	return &entity.Service{
		ID:             objectID,
		ServiceGroupID: serviceGroupID,
		Name:           req.Name,
		IsActive:       req.IsActive,
		Description:    req.Description,
	}, nil
}

