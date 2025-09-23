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

type ServiceGroupRepository interface {
	Upload(ctx context.Context, group *model.ServiceGroup) error
	GetAll(ctx context.Context) ([]*model.ServiceGroup, error)
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

func (r *serviceGroupRepository) GetAll(ctx context.Context) ([]*model.ServiceGroup, error) {
	cursor, err := r.collection.Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{"order", 1}}))
	if err != nil {
		return nil, err
	}
	var groups []*model.ServiceGroup
	if err := cursor.All(ctx, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}
