package repository

import (
	"context"
	"time"

	"github.com/senbox/services-management/internal/domain/entity"
	"github.com/senbox/services-management/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type serviceRepositoryMongo struct {
	collection *mongo.Collection
}

func NewServiceRepositoryMongo(db *mongo.Database) repository.ServiceRepository {
	return &serviceRepositoryMongo{
		collection: db.Collection("services"),
	}
}

func (r *serviceRepositoryMongo) Create(ctx context.Context, service *entity.Service) error {
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()
	
	result, err := r.collection.InsertOne(ctx, service)
	if err != nil {
		return err
	}
	
	service.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *serviceRepositoryMongo) FindByID(ctx context.Context, id primitive.ObjectID) (*entity.Service, error) {
	var service entity.Service
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&service)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &service, nil
}

func (r *serviceRepositoryMongo) FindAll(ctx context.Context) ([]*entity.Service, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var services []*entity.Service
	if err := cursor.All(ctx, &services); err != nil {
		return nil, err
	}

	return services, nil
}

func (r *serviceRepositoryMongo) FindByServiceGroupID(ctx context.Context, serviceGroupID primitive.ObjectID) ([]*entity.Service, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"service_group_id": serviceGroupID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var services []*entity.Service
	if err := cursor.All(ctx, &services); err != nil {
		return nil, err
	}

	return services, nil
}

func (r *serviceRepositoryMongo) Update(ctx context.Context, service *entity.Service) error {
	service.UpdatedAt = time.Now()
	
	filter := bson.M{"_id": service.ID}
	update := bson.M{
		"$set": bson.M{
			"service_group_id": service.ServiceGroupID,
			"name":             service.Name,
			"is_active":        service.IsActive,
			"description":      service.Description,
			"updated_at":       service.UpdatedAt,
		},
	}
	
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *serviceRepositoryMongo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

