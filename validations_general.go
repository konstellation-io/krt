package main

import (
	"regexp"

	"github.com/konstellation-io/krt/errors"
)

const maxFieldNameLength = 20

func isValidResourceName(name string) bool {
	reResourceName := regexp.MustCompile("^[a-z0-9]([-a-z0-9]*[a-z0-9])?$")
	return reResourceName.MatchString(name)
}

func validateName(name, nameLocation string) error {
	if name == "" {
		return errors.MissingRequiredFieldError(nameLocation)
	}
	if !isValidResourceName(name) {
		return errors.InvalidFieldNameError(nameLocation)
	}
	if len(name) > maxFieldNameLength {
		return errors.InvalidLengthFieldError(nameLocation, maxFieldNameLength)
	}
	return nil
}
