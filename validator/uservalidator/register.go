package uservalidator

import (
	"regexp"

	"github.com/aghaghiamh/ava/domain"
	"github.com/aghaghiamh/ava/pkg/richerr"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v UserValidator) ValidateRegisterRequest(req domain.RegisterRequest) (map[string]string, error) {
	const op = "validator.ValidateRegisterRequest"

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 20)),

		validation.Field(
			&req.PhoneNumber,
			validation.Match(regexp.MustCompile(`^(\(?\+98\)?)?[-\s]?(09)(\d{9})$`)).Error("Phone number does not satisfy the valid pattern of `(+98) 09xxxxxxxxx`.")),
	)

	if err != nil {
		fieldErrors := map[string]string{}
		if vErr, ok := err.(validation.Errors); ok {
			for key, val := range vErr {
				if val != nil {
					fieldErrors[key] = val.Error()
				}
			}
		}
		return fieldErrors, richerr.New(op).WithError(err).WithCode(richerr.ErrInvalidInput)
	}

	return nil, nil
}
