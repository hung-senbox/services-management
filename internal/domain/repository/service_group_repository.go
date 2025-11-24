package repository

import (
	"context"

	"services-management/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceGroupRepository interface {
	Create(ctx context.Context, serviceGroup *entity.ServiceGroup) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*entity.ServiceGroup, error)
	FindAll(ctx context.Context) ([]*entity.ServiceGroup, error)
	Update(ctx context.Context, serviceGroup *entity.ServiceGroup) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
