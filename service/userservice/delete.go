package userservice

import (
	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/pkg/richerr"
)

func (s *Service) Delete(req domain.DelAccountRequest) (domain.DelAccountResponse, error) {
	const op = "userservice.DelAccount"
	gErr := s.repo.DelByID(req.UserID)
	if gErr != nil {

		return domain.DelAccountResponse{}, richerr.New(op).WithError(gErr)
	}

	return domain.DelAccountResponse{}, nil
}
