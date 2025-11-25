package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceGroup struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Order       int                `bson:"order" json:"order"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	Description string             `bson:"description" json:"description"`
	Url         string             `bson:"url" json:"url"`
	Icon        string             `bson:"icon" json:"icon"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
