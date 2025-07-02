package userhandler

import (
	"github.com/aghaghiamh/ava/service/userservice"
	"github.com/aghaghiamh/ava/validator/uservalidator"
)

type HandlerConfig struct {
	DefaultPageSizeStr string `mapstructure:"default_page_size"`
	DefaultMaxPageSize int    `mapstructure:"default_max_page_size"`
}

type Handler struct {
	config    HandlerConfig
	validator uservalidator.UserValidator
	userSvc   userservice.Service
}

func New(config HandlerConfig, validator uservalidator.UserValidator, userSvc userservice.Service) Handler {
	return Handler{
		config:    config,
		validator: validator,
		userSvc:   userSvc,
	}
}
