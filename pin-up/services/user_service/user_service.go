package user_service

import (
	"go-app/pin-up/domain/user"
	"go-app/pin-up/repositories/user_repo"
)

type UserServiceInterface interface {
	CreateUser(user *user.User) (*user.User, error)
}

type userService struct {
	Repo user_repo.Repo
}

func NewUserService(repo user_repo.Repo) UserServiceInterface  {
	return &userService{
		Repo: repo,
	}
}

func (u *userService) CreateUser(user *user.User) (*user.User, error)  {
	return u.Repo.CreateUser(user)
}