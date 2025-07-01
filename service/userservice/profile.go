package userservice

import (
	"context"
	"fmt"

	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/pkg/richerr"
)

func (s *Service) GetProfile(ctx context.Context, req domain.ProfileRequest) (domain.ProfileResponse, error) {
	const op = "userservice.GetUserProfile"
	user, gErr := s.repo.GetUserByID(ctx, req.UserID)
	if gErr != nil {

		return domain.ProfileResponse{}, richerr.New(op).
			WithError(gErr).
			WithMessage(fmt.Sprintf("User with %d id does not exist. Please register first.", req.UserID))
	}

	return domain.ProfileResponse{
		UserInfo: domain.UserInfo{
			UserID:      user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
