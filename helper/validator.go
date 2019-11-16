package helper

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

// Validator :
type Validator struct {
	validator *govalidator.Validator
}

// Validate :
func (v *Validator) Validate(i interface{}, opts govalidator.Options) url.Values {
	v.validator.Opts = opts
	return v.validator.Validate()
}
