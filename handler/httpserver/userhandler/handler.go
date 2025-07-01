package userhandler

import (
	"github.com/aghaghiamh/ava/service/userservice"
	"github.com/aghaghiamh/ava/validator/uservalidator"
)

type Handler struct {
	userSvc    userservice.Service
	validator  uservalidator.UserValidator
}

func New(userSvc userservice.Service, validator uservalidator.UserValidator) Handler {
	return Handler{
		userSvc:    userSvc,
		validator:  validator,
	}
}
