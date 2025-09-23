package repository

import (
	"context"
	"services-management/internal/sv_management/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceRepository interface {
	Upload(ctx context.Context, service *model.Service) error
	GetAll(ctx context.Context) ([]*model.Service, error)
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

func (r *serviceRepository) GetAll(ctx context.Context) ([]*model.Service, error) {
	cursor, err := r.collection.Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{"order", 1}}))
	if err != nil {
		return nil, err
	}
	var services []*model.Service
	if err := cursor.All(ctx, &services); err != nil {
		return nil, err
	}
	return services, nil
}
