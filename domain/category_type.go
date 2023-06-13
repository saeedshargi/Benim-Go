package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryType struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Title          string             `bson:"title" json:"title" validate:"required"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
	IsDeleted      bool               `bson:"isDeleted" json:"-"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"-"`
	LastModifiedAt primitive.DateTime `bson:"lastModifiedAt" json:"-"`
}
