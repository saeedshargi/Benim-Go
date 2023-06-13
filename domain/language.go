package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Language struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Name           string             `bson:"name" json:"name" validate:"required"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
	IsDeleted      bool               `bson:"isDeleted" json:"-"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"-"`
	LastModifiedAt primitive.DateTime `bson:"lastModifiedAt" json:"-"`
}
