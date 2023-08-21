package repository

import (
	"Benim/domain"
	"Benim/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type languageRepository struct {
	database   mongo.Database
	collection string
}

func NewLanguageRepository(db mongo.Database, collection string) domain.LanguageRepository {
	return &languageRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.LanguageRepository.
func (b *languageRepository) Create(ctx context.Context, language *domain.Language) error {
	collection := b.database.Collection(b.collection)

	_, err := collection.InsertOne(ctx, language)

	return err
}

// Delete implements domain.LanguageRepository.
func (l *languageRepository) Delete(ctx context.Context, id int) error {
	collection := l.database.Collection(l.collection)

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})

	return err
}

// GetAll implements domain.LanguageRepository.
func (l *languageRepository) GetAll(ctx context.Context) ([]domain.Language, error) {
	collection := l.database.Collection(l.collection)

	var languages []domain.Language

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &languages)
	if err != nil {
		return nil, err
	}

	return languages, nil
}

// GetById implements domain.LanguageRepository.
func (*languageRepository) GetById(ctx context.Context, id int) (domain.Language, error) {
	panic("unimplemented")
}

// Update implements domain.LanguageRepository.
func (*languageRepository) Update(ctx context.Context, id int, language *domain.Language) error {
	panic("unimplemented")
}
