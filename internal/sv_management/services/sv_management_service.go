package service

import (
	"context"
	"services-management/internal/sv_management/dto/request"
	"services-management/internal/sv_management/model"
	"services-management/internal/sv_management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SvManagementService interface {
	UploadService(ctx context.Context, req request.UploadServiceRequest) error
}

type svManagementService struct {
	repository repository.ServiceRepository
}

func NewSvManagementService(repository repository.ServiceRepository) SvManagementService {
	return &svManagementService{
		repository: repository,
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
	return s.repository.Upload(ctx, service)
}
