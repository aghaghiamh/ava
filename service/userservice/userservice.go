package userservice

import (
	"context"

	"github.com/aghaghiamh/ava/entity"
)

type UserRepo interface {
	Register(user entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint) (entity.User, error)
	DelByID(userID uint) error
	ListWithPagination(ctx context.Context, page, pageSize int) ([]entity.User, error)
}

type Service struct {
	repo UserRepo
}

func New(repo UserRepo) Service {
	return Service{
		repo: repo,
	}
}
