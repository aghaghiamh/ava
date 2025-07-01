package userservice

import (
	"context"

	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/pkg/richerr"
)

func (s *Service) ListWithPagination(ctx context.Context, req domain.ListRequest) (domain.ListResponse, error) {
	const op = "userservice.ListWithPagination"
	users, lErr := s.repo.ListWithPagination(ctx, req.Page, req.PageSize)
	if lErr != nil {

		return domain.ListResponse{}, richerr.New(op).
			WithError(lErr)
	}

	usersInfo := []domain.UserInfo{}
	for _, ui := range users {
		usersInfo = append(usersInfo, domain.UserInfo{
			UserID:      ui.ID,
			Name:        ui.Name,
			PhoneNumber: ui.PhoneNumber,
		})
	}

	return domain.ListResponse{
		UsersInfo: usersInfo,
	}, nil
}
