package service

import (
	"context"
	"services-management/internal/sv_management/dto/request"
	"services-management/internal/sv_management/dto/response"
	"services-management/internal/sv_management/mapper"
	"services-management/internal/sv_management/model"
	"services-management/internal/sv_management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SvManagementService interface {
	UploadService(ctx context.Context, req request.UploadServiceRequest) error
	GetServices(ctx context.Context) ([]*response.ServicesResponse, error)
}

type svManagementService struct {
	serviceRepo      repository.ServiceRepository
	serviceGroupRepo repository.ServiceGroupRepository
}

func NewSvManagementService(
	serviceRepo repository.ServiceRepository,
	serviceGroupRepo repository.ServiceGroupRepository,
) *svManagementService {
	return &svManagementService{
		serviceRepo:      serviceRepo,
		serviceGroupRepo: serviceGroupRepo,
	}
}

func (s *svManagementService) UploadService(ctx context.Context, req request.UploadServiceRequest) error {

	service := &model.Service{
		ID:      primitive.NewObjectID(),
		Title:   req.Title,
		Url:     req.Url,
		Order:   req.Order,
		GroupID: req.GroupID,
	}
	return s.serviceRepo.Upload(ctx, service)
}

func (s *svManagementService) GetServices(ctx context.Context) ([]*response.ServicesResponse, error) {
	// Lấy groups
	groups, err := s.serviceGroupRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Lấy services
	services, err := s.serviceRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.MapServicesResponse(groups, services), nil
}
