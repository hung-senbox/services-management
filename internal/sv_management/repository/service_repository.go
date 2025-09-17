package repository

import (
	"context"
	"services-management/internal/sv_management/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRepository interface {
	Upload(ctx context.Context, service *model.Service) error
}

type serviceRepository struct {
	collection *mongo.Collection
}

func NewServiceRepository(collection *mongo.Collection) ServiceRepository {
	return &serviceRepository{
		collection: collection,
	}
}

func (r *serviceRepository) Upload(ctx context.Context, service *model.Service) error {
	// Nếu chưa có _id thì tự sinh
	if service.ID.IsZero() {
		service.ID = primitive.NewObjectID()
	}

	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, service)
	return err
}
