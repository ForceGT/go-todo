package validator

import val "github.com/go-playground/validator"

type EchoRequestValidator struct {
	validator *val.Validate
}

func NewEchoRequestValidator() *EchoRequestValidator {
	return &EchoRequestValidator{
		validator: val.New(),
	}
}

func (v *EchoRequestValidator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err != nil {
		// TODO: Handle error array
		// for _, err := range err.(val.ValidationErrors) {
		//
		//}
		return err
	}
	return nil
}
