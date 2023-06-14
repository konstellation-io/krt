package krt

import (
	"regexp"

	"github.com/konstellation-io/krt/pkg/errors"
)

const MaxFieldNameLength = 20

func isValidResourceName(name string) bool {
	reResourceName := regexp.MustCompile(`^[a-z0-9]([-|.a-z0-9]*[a-z0-9])?$`)
	return reResourceName.MatchString(name)
}

func validateName(name, nameLocation string) error {
	if name == "" {
		return errors.MissingRequiredFieldError(nameLocation)
	}

	if !isValidResourceName(name) {
		return errors.InvalidFieldNameError(nameLocation)
	}

	if len(name) > MaxFieldNameLength {
		return errors.InvalidLengthFieldError(nameLocation, MaxFieldNameLength)
	}

	return nil
}
