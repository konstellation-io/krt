package krt

import (
	"regexp"

	"github.com/konstellation-io/krt/pkg/errors"
)

const MaxFieldNameLength = 20

func isValidVersion(version string) bool {
	reVersion := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	return reVersion.MatchString(version)
}

func isValidResourceName(name string) bool {
	reResourceName := regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	return reResourceName.MatchString(name)
}

func validateVersion(version, versionLocation string) error {
	if version == "" {
		return errors.MissingRequiredFieldError(versionLocation)
	}

	if !isValidVersion(version) {
		return errors.InvalidVersionTagError(versionLocation)
	}

	return nil
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
