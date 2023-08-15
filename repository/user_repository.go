package repository

import (
	"Benim/domain"
	"Benim/internal/tokenutil"
	"Benim/mongo"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

// GetAll implements domain.UserRepository.
func (u *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	collection := u.database.Collection(u.collection)
	var users []domain.User

	results, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var user domain.User
		if err = results.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByEmail implements domain.UserRepository.
func (u *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// Login implements domain.UserRepository.
func (u *userRepository) Login(ctx context.Context, email string, password string, secret string) (domain.UserResponse, error) {
	collection := u.database.Collection(u.collection)

	user, err := u.GetByEmail(ctx, email)
	if err != nil {
		return domain.UserResponse{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.UserResponse{}, errors.New("password is wrong")
	}

	updateUser := bson.M{"lastLoginAt": time.Now().UTC()}

	result, err := collection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": updateUser})
	if err != nil {
		return domain.UserResponse{}, errors.New("login failed with unexpected error")
	}

	var updatedUser domain.User
	if result.MatchedCount == 1 {
		err := collection.FindOne(ctx, bson.M{"id": user.ID}).Decode(&updatedUser)

		if err != nil {
			return domain.UserResponse{}, errors.New("login failed with unexpected error")
		}
	}

	accessToken, err := tokenutil.CreateAccessToken(&updatedUser, secret, 20)
	if err != nil {
		return domain.UserResponse{}, err
	}

	refreshToken, err := tokenutil.CreateRefreshToken(&updatedUser, secret, 1)
	if err != nil {
		return domain.UserResponse{}, err
	}

	userResponse := &domain.UserResponse{
		ID:           updatedUser.ID.Hex(),
		Email:        updatedUser.Email,
		FullName:     updatedUser.FullName,
		AccessToekn:  accessToken,
		RefreshToken: refreshToken,
	}

	return *userResponse, nil
}

// Register implements domain.UserRepository.
func (u *userRepository) Register(ctx context.Context, user *domain.User, secret string) (domain.UserResponse, error) {
	collection := u.database.Collection(u.collection)

	existUser, err := u.GetByEmail(ctx, user.Email)
	if err != nil {
		return domain.UserResponse{}, err
	}

	if existUser != (domain.User{}) {
		return domain.UserResponse{}, errors.New("user with this email already exit")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return domain.UserResponse{}, errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	_, err = collection.InsertOne(ctx, user)

	accessToken, err := tokenutil.CreateAccessToken(user, secret, 20)
	if err != nil {
		return domain.UserResponse{}, err
	}

	refreshToken, err := tokenutil.CreateRefreshToken(user, secret, 1)
	if err != nil {
		return domain.UserResponse{}, err
	}

	userResponse := &domain.UserResponse{
		ID:           user.ID.Hex(),
		Email:        user.Email,
		FullName:     user.FullName,
		AccessToekn:  accessToken,
		RefreshToken: refreshToken,
	}

	return *userResponse, nil
}
