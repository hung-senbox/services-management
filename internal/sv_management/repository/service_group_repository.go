package repository

import (
	"context"
	"services-management/internal/sv_management/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceGroupRepository interface {
	Upload(ctx context.Context, group *model.ServiceGroup) error
}

type serviceGroupRepository struct {
	collection *mongo.Collection
}

func NewServiceGroupRepository(collection *mongo.Collection) ServiceGroupRepository {
	return &serviceGroupRepository{
		collection: collection,
	}
}

func (r *serviceGroupRepository) Upload(ctx context.Context, group *model.ServiceGroup) error {
	// Nếu chưa có _id thì tự sinh
	if group.ID.IsZero() {
		group.ID = primitive.NewObjectID()
	}

	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, group)
	return err
}
