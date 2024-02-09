package kubeutil

import (
	"errors"

	"k8s.io/apimachinery/pkg/util/validation"
)

func ValidateNodeSelectorKey(value string) error {
	var errs error

	for _, msg := range validation.IsQualifiedName(value) {
		errs = errors.Join(errs, errors.New(msg))
	}

	return errs
}

func ValidateNodeSelectorValue(value string) error {
	var errs error

	for _, msg := range validation.IsValidLabelValue(value) {
		errs = errors.Join(errs, errors.New(msg))
	}

	return errs
}
