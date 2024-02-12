package kubeutil

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	_maxNameLength         = 63
	_maxValueLength        = 63
	_maxDNSSubdomainLength = 253

	_alphaNumFmt     = "[A-Za-z0-9]"
	_alphaNumWithFmt = "[-A-Za-z0-9_.]"
	_validDNSFmt     = "[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"

	_qualifiedNameFmt = "(" + _alphaNumFmt + _alphaNumWithFmt + "*)?" + _alphaNumFmt
)

var (
	ErrInvalidKeyPrefix = errors.New("invalid key prefix")
	ErrInvalidKeyFormat = errors.New("invalid key format: key must be a valid name with an optional DNS subdomain prefix joined by '/'")
	ErrInvalidKeyName   = errors.New("invalid key name")
	ErrInvalidValue     = errors.New("invalid value")

	_validQualifiedNameRegexp = regexp.MustCompile("^" + _qualifiedNameFmt + "$")
	_validDNSSubdomainRegexp  = regexp.MustCompile("^" + _validDNSFmt + "(\\." + _validDNSFmt + ")*$")
)

func ValidateNodeSelectorKey(value string) error {
	keySections := strings.Split(value, "/")

	switch len(keySections) {
	case 1:
		return validateKeyName(keySections[0])
	case 2:
		prefix, name := keySections[0], keySections[1]

		if err := validateKeyPrefix(prefix); err != nil {
			return fmt.Errorf("invalid prefix %q: %w", prefix, err)
		}

		if err := validateKeyName(name); err != nil {
			return fmt.Errorf("invalid name %q: %w", name, err)
		}

	default:
		return ErrInvalidKeyFormat
	}

	return nil
}

func validateKeyPrefix(prefix string) error {
	if prefix == "" {
		return fmt.Errorf("%w: prefix cannot be empty", ErrInvalidKeyPrefix)
	}

	if len(prefix) > _maxDNSSubdomainLength {
		return fmt.Errorf("%w: prefix too long", ErrInvalidKeyPrefix)
	}

	if !_validDNSSubdomainRegexp.MatchString(prefix) {
		return fmt.Errorf("%w: key prefix must match the regexp %q",
			ErrInvalidKeyPrefix,
			_validDNSFmt,
		)
	}

	return nil
}

func validateKeyName(name string) error {
	if name == "" {
		return fmt.Errorf("%w: missing mandatory key name", ErrInvalidKeyName)
	}

	if len(name) > _maxNameLength {
		return fmt.Errorf(
			"%w: name length must be less than %d characters",
			ErrInvalidKeyName,
			_maxNameLength,
		)
	}

	if !_validQualifiedNameRegexp.MatchString(name) {
		return fmt.Errorf("%w: name must match the regex %q",
			ErrInvalidKeyName,
			_validQualifiedNameRegexp,
		)
	}

	return nil
}

func ValidateNodeSelectorValue(value string) error {
	if len(value) > _maxValueLength {
		return fmt.Errorf(
			"%w: value length must be less than %d characters",
			ErrInvalidValue,
			_maxValueLength,
		)
	}

	if !_validQualifiedNameRegexp.MatchString(value) {
		return fmt.Errorf(
			"%w: value must match the regex %q",
			ErrInvalidValue,
			_validQualifiedNameRegexp,
		)
	}

	return nil
}
