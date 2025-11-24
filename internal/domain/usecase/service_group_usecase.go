package usecase

import (
	"context"

	"services-management/internal/domain/entity"
	"services-management/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceGroupUseCase interface {
	CreateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error
	GetServiceGroupByID(ctx context.Context, id string) (*entity.ServiceGroup, error)
	GetAllServiceGroups(ctx context.Context) ([]*entity.ServiceGroup, error)
	UpdateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error
	DeleteServiceGroup(ctx context.Context, id string) error
}

type serviceGroupUseCase struct {
	serviceGroupRepo repository.ServiceGroupRepository
}

func NewServiceGroupUseCase(serviceGroupRepo repository.ServiceGroupRepository) ServiceGroupUseCase {
	return &serviceGroupUseCase{
		serviceGroupRepo: serviceGroupRepo,
	}
}

func (uc *serviceGroupUseCase) CreateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error {
	return uc.serviceGroupRepo.Create(ctx, serviceGroup)
}

func (uc *serviceGroupUseCase) GetServiceGroupByID(ctx context.Context, id string) (*entity.ServiceGroup, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return uc.serviceGroupRepo.FindByID(ctx, objectID)
}

func (uc *serviceGroupUseCase) GetAllServiceGroups(ctx context.Context) ([]*entity.ServiceGroup, error) {
	return uc.serviceGroupRepo.FindAll(ctx)
}

func (uc *serviceGroupUseCase) UpdateServiceGroup(ctx context.Context, serviceGroup *entity.ServiceGroup) error {
	return uc.serviceGroupRepo.Update(ctx, serviceGroup)
}

func (uc *serviceGroupUseCase) DeleteServiceGroup(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return uc.serviceGroupRepo.Delete(ctx, objectID)
}
