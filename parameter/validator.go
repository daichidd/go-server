package parameter

import (
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

var valid *validator.Validate

func init() {
	valid = validator.New()
}

func ValidParam(v interface{}) error {
	if err := valid.Struct(v); err != nil {
		return errors.Wrap(err, "valid.Struct failed")
	}
	return nil
}
