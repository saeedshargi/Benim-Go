package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBook = "books"
)

type Book struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title          string             `bson:"title" json:"title" validate:"required"`
	Category       []Category         `bson:"category" json:"category"`
	Language       []Language         `bson:"language" json:"language"`
	Author         string             `bson:"author" json:"author" validate:"required"`
	Publisher      string             `bson:"publisher" json:"publisher" validate:"required"`
	PublishYear    int                `bson:"publishYear" json:"publishYear" validate:"required"`
	ISBN           string             `bson:"isbn" json:"isbn"`
	PageCount      int                `bson:"pageCount" json:"pageCount"`
	Summary        string             `bson:"summary" json:"summary"`
	IsFavorite     bool               `bson:"isFavorite" json:"isFavorite"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
	IsDeleted      bool               `bson:"isDeleted" json:"-"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"-"`
	LastModifiedAt primitive.DateTime `bson:"lastModifiedAt" json:"-"`
}

type BookRepository interface {
	Create(ctx context.Context, book *Book) error
	Update(ctx context.Context, id int, book *Book) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
}

type BookUsecase interface {
	Create(ctx context.Context, book *Book) error
	Update(ctx context.Context, id int, book *Book) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
}
