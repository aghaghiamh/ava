package uservalidator

import (
	"context"

	"github.com/aghaghiamh/ava/entity"
)

const (
	PhoneNumberRegex = `^(\(?\+98\)?)?[-\s]?(09)(\d{9})$`
)

type UserRepo interface {
	GetUserByID(ctx context.Context, user_id uint) (entity.User, error)
}

type UserValidator struct {
	repo UserRepo
}

func New(repo UserRepo) UserValidator {
	return UserValidator{
		repo: repo,
	}
}
