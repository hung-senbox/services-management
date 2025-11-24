package usecase

import (
	"context"

	"github.com/senbox/services-management/internal/domain/entity"
	"github.com/senbox/services-management/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceUseCase interface {
	CreateService(ctx context.Context, service *entity.Service) error
	GetServiceByID(ctx context.Context, id string) (*entity.Service, error)
	GetAllServices(ctx context.Context) ([]*entity.Service, error)
	GetServicesByGroupID(ctx context.Context, serviceGroupID string) ([]*entity.Service, error)
	UpdateService(ctx context.Context, service *entity.Service) error
	DeleteService(ctx context.Context, id string) error
}

type serviceUseCase struct {
	serviceRepo repository.ServiceRepository
}

func NewServiceUseCase(serviceRepo repository.ServiceRepository) ServiceUseCase {
	return &serviceUseCase{
		serviceRepo: serviceRepo,
	}
}

func (uc *serviceUseCase) CreateService(ctx context.Context, service *entity.Service) error {
	return uc.serviceRepo.Create(ctx, service)
}

func (uc *serviceUseCase) GetServiceByID(ctx context.Context, id string) (*entity.Service, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return uc.serviceRepo.FindByID(ctx, objectID)
}

func (uc *serviceUseCase) GetAllServices(ctx context.Context) ([]*entity.Service, error) {
	return uc.serviceRepo.FindAll(ctx)
}

func (uc *serviceUseCase) GetServicesByGroupID(ctx context.Context, serviceGroupID string) ([]*entity.Service, error) {
	objectID, err := primitive.ObjectIDFromHex(serviceGroupID)
	if err != nil {
		return nil, err
	}
	return uc.serviceRepo.FindByServiceGroupID(ctx, objectID)
}

func (uc *serviceUseCase) UpdateService(ctx context.Context, service *entity.Service) error {
	return uc.serviceRepo.Update(ctx, service)
}

func (uc *serviceUseCase) DeleteService(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return uc.serviceRepo.Delete(ctx, objectID)
}

