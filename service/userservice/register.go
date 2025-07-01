package userservice

import (
	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/entity"
	"github.com/aghaghiamh/ava/pkg/richerr"
)

func (s *Service) Register(req domain.RegisterRequest) (domain.RegisterResponse, error) {
	const op = "userservice.Register"

	user := entity.User{
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	var regErr error
	user, regErr = s.repo.Register(user)
	if regErr != nil {
		return domain.RegisterResponse{}, richerr.New(op).WithError(regErr)
	}

	return domain.RegisterResponse{
		UserInfo: domain.UserInfo{
			UserID:      user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
