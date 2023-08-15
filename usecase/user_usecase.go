package usecase

import (
	"Benim/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contexttimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contexttimeout: timeout,
	}
}

// GetAll implements domain.UserUsecase.
func (u *userUsecase) GetAll(ctx context.Context) ([]domain.User, error) {
	return u.userRepository.GetAll(ctx)
}

// GetByEmail implements domain.UserUsecase.
func (u *userUsecase) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.userRepository.GetByEmail(ctx, email)
}

// Login implements domain.UserUsecase.
func (u *userUsecase) Login(ctx context.Context, email string, password string, secret string) (domain.UserResponse, error) {
	return u.userRepository.Login(ctx, email, password, secret)
}

// Register implements domain.UserUsecase.
func (u *userUsecase) Register(ctx context.Context, user *domain.SignupUser, secrest string) (domain.UserResponse, error) {
	newUser := &domain.User{
		ID:             primitive.NewObjectID(),
		Email:          user.Email,
		Password:       user.Password,
		FullName:       user.FullName,
		LastLoginAt:    0,
		IsActive:       true,
		IsDeleted:      false,
		CreatedAt:      primitive.DateTime(time.Now().UTC().Unix()),
		LastModifiedAt: 0,
	}

	return u.userRepository.Register(ctx, newUser, secrest)
}
