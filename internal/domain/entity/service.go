package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ServiceGroupID primitive.ObjectID `bson:"service_group_id" json:"service_group_id"`
	Name           string             `bson:"name" json:"name"`
	IsActive       bool               `bson:"is_active" json:"is_active"`
	Description    string             `bson:"description" json:"description"`
	Url            string             `bson:"url" json:"url"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}
