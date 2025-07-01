package userservice

import (
	"context"

	"github.com/aghaghiamh/ava/entity"
)

type UserRepo interface {
	Register(user entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, user_id uint) (entity.User, error)
	DelByID(user_id uint) error
}

type Service struct {
	repo UserRepo
}

func New(repo UserRepo) Service {
	return Service{
		repo: repo,
	}
}
