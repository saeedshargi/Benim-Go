package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionLanguage = "languages"
)

type Language struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name" validate:"required"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
	IsDeleted      bool               `bson:"isDeleted" json:"-"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"-"`
	LastModifiedAt primitive.DateTime `bson:"lastModifiedAt" json:"-"`
}

type LanguageRepository interface {
	Create(ctx context.Context, language *Language) error
	Update(ctx context.Context, id int, language *Language) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Language, error)
	GetAll(ctx context.Context) ([]Language, error)
}

type LanguageUsecase interface {
	Create(ctx context.Context, language *Language) error
	Update(ctx context.Context, id int, language *Language) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Language, error)
	GetAll(ctx context.Context) ([]Language, error)
}
