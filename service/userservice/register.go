package userservice

import (
	"github.com/aghaghiamh/ava/dto"
	"github.com/aghaghiamh/ava/entity"
	"github.com/aghaghiamh/ava/pkg/richerr"
)

func (s *Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	const op = "userservice.Register"

	user := entity.User{
		PhoneNumber:    req.PhoneNumber,
		Name:           req.Name,
	}
	
	var regErr error
	user, regErr = s.repo.Register(user)
	if regErr != nil {
		return dto.RegisterResponse{}, richerr.New(op).WithError(regErr)
	}

	return dto.RegisterResponse{
		UserInfo: dto.UserInfo{
			UserID:      user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
