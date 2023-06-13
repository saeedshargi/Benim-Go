package repository

import (
	"Benim/domain"
	"Benim/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type bookRepository struct {
	database   mongo.Database
	collection string
}

func NewBookRepository(db mongo.Database, collection string) domain.BookRepository {
	return &bookRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.BookRepository.
func (b *bookRepository) Create(ctx context.Context, book *domain.Book) error {
	collection := b.database.Collection(b.collection)

	_, err := collection.InsertOne(ctx, book)

	return err
}

// Delete implements domain.BookRepository.
func (b *bookRepository) Delete(ctx context.Context, id int) error {
	collection := b.database.Collection(b.collection)

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})

	return err
}

// GetAll implements domain.BookRepository.
func (*bookRepository) GetAll(ctx context.Context) ([]domain.Book, error) {
	panic("unimplemented")
}

// GetById implements domain.BookRepository.
func (*bookRepository) GetById(ctx context.Context, id int) (domain.Book, error) {
	panic("unimplemented")
}

// Update implements domain.BookRepository.
func (*bookRepository) Update(ctx context.Context, id int, book *domain.Book) error {
	panic("unimplemented")
}
