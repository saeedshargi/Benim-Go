package usecase

import (
	"Benim/domain"
	"context"
	"time"
)

type languageUsecase struct {
	languageRepository domain.LanguageRepository
	cotextTimeout      time.Duration
}

func NewLanguageUsecase(languageRepository domain.LanguageRepository, timeout time.Duration) domain.LanguageUsecase {
	return &languageUsecase{
		languageRepository: languageRepository,
		cotextTimeout:      timeout,
	}
}

// Create implements domain.LanguageUsecase.
func (bu *languageUsecase) Create(ctx context.Context, language *domain.Language) error {
	return bu.languageRepository.Create(ctx, language)
}

// Delete implements domain.LanguageUsecase.
func (*languageUsecase) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetAll implements domain.LanguageUsecase.
func (bu *languageUsecase) GetAll(ctx context.Context) ([]domain.Language, error) {
	return bu.languageRepository.GetAll(ctx)
}

// GetById implements domain.LanguageUsecase.
func (*languageUsecase) GetById(ctx context.Context, id int) (domain.Language, error) {
	panic("unimplemented")
}

// Update implements domain.LanguageUsecase.
func (*languageUsecase) Update(ctx context.Context, id int, language *domain.Language) error {
	panic("unimplemented")
}
