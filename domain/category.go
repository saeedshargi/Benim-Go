package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Title          string             `bson:"title" json:"title" validate:"required"`
	CategoryType   []CategoryType     `bson:"categoryType" json:"categoryType"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
	IsDeleted      bool               `bson:"isDeleted" json:"-"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"-"`
	LastModifiedAt primitive.DateTime `bson:"lastModifiedAt" json:"-"`
}
