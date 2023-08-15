package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUsers = "users"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Email          string             `bson:"email" json:"email" validate:"required"`
	Password       string             `bson:"password" json:"password" validate:"required"`
	FullName       string             `bson:"fullName" json:"fullName"`
	LastLoginAt    primitive.DateTime `bson:"lastLoginAt" json:"-"`
	IsActive       bool               `bson:"isActive" json:"isActive"`
	IsDeleted      bool               `bson:"isDeleted" json:"-"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"-"`
	LastModifiedAt primitive.DateTime `bson:"lastModifiedAt" json:"-"`
}

type UserResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
	AccessToekn  string `json:"accessToken"`
	RefreshToken string `json:"refreshTken"`
}

type SignupUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"fullName"`
}

type UserRepository interface {
	Register(ctx context.Context, user *User, secrest string) (UserResponse, error)
	Login(ctx context.Context, email string, password string, secret string) (UserResponse, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
}

type UserUsecase interface {
	Register(ctx context.Context, user *SignupUser, secrest string) (UserResponse, error)
	Login(ctx context.Context, email string, password string, secret string) (UserResponse, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
}
