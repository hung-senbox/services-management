package service

import (
	"context"
	"services-management/internal/sv_management/dto/request"
	"services-management/internal/sv_management/model"
	"services-management/internal/sv_management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SVGroupService interface {
	UploadServiceGroup(ctx context.Context, req request.UploadServiceGroupRequest) error
}

type svGroupService struct {
	repository repository.ServiceGroupRepository
}

func NewSVGroupService(repository repository.ServiceGroupRepository) SVGroupService {
	return &svGroupService{
		repository: repository,
	}
}

func (s *svGroupService) UploadServiceGroup(ctx context.Context, req request.UploadServiceGroupRequest) error {

	serviceGroup := &model.ServiceGroup{
		ID:    primitive.NewObjectID(),
		Title: req.Title,
		Order: req.Order,
	}
	return s.repository.Upload(ctx, serviceGroup)
}
