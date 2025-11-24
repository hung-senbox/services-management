package repository

import (
	"context"

	"github.com/senbox/services-management/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceRepository interface {
	Create(ctx context.Context, service *entity.Service) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*entity.Service, error)
	FindAll(ctx context.Context) ([]*entity.Service, error)
	FindByServiceGroupID(ctx context.Context, serviceGroupID primitive.ObjectID) ([]*entity.Service, error)
	Update(ctx context.Context, service *entity.Service) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

