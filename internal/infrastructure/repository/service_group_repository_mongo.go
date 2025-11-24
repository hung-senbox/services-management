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

type serviceGroupRepositoryMongo struct {
	collection *mongo.Collection
}

func NewServiceGroupRepositoryMongo(db *mongo.Database) repository.ServiceGroupRepository {
	return &serviceGroupRepositoryMongo{
		collection: db.Collection("service_groups"),
	}
}

func (r *serviceGroupRepositoryMongo) Create(ctx context.Context, serviceGroup *entity.ServiceGroup) error {
	serviceGroup.CreatedAt = time.Now()
	serviceGroup.UpdatedAt = time.Now()
	
	result, err := r.collection.InsertOne(ctx, serviceGroup)
	if err != nil {
		return err
	}
	
	serviceGroup.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *serviceGroupRepositoryMongo) FindByID(ctx context.Context, id primitive.ObjectID) (*entity.ServiceGroup, error) {
	var serviceGroup entity.ServiceGroup
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&serviceGroup)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &serviceGroup, nil
}

func (r *serviceGroupRepositoryMongo) FindAll(ctx context.Context) ([]*entity.ServiceGroup, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var serviceGroups []*entity.ServiceGroup
	if err := cursor.All(ctx, &serviceGroups); err != nil {
		return nil, err
	}

	return serviceGroups, nil
}

func (r *serviceGroupRepositoryMongo) Update(ctx context.Context, serviceGroup *entity.ServiceGroup) error {
	serviceGroup.UpdatedAt = time.Now()
	
	filter := bson.M{"_id": serviceGroup.ID}
	update := bson.M{
		"$set": bson.M{
			"name":        serviceGroup.Name,
			"order":       serviceGroup.Order,
			"is_active":   serviceGroup.IsActive,
			"description": serviceGroup.Description,
			"icon_key":    serviceGroup.IconKey,
			"updated_at":  serviceGroup.UpdatedAt,
		},
	}
	
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *serviceGroupRepositoryMongo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

