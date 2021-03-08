package user_repo

import (
	"context"
	"go-app/pin-up/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repo interface {
	CreateUser(user *user.User) (*user.User, error)
}

type UserRepo struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) Repo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(user *user.User) (*user.User, error)   {
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)

	collection := u.db.Database("pin-up").Collection("users")
	_, err := collection.InsertOne(ctx, *user)

	if err != nil {
		panic(err)
	}
	return user, nil
}