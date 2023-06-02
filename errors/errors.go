package errors

import (
	"errors"
	"fmt"
)

func MergeErrors(err1, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	}
	if err1 == nil && err2 != nil {
		return err2
	}
	if err1 != nil && err2 == nil {
		return err1
	}

	return fmt.Errorf("%w\n%w", err1, err2)
}

var ErrMissingRequiredField = errors.New("missing required field")
var ErrInvalidFieldName = errors.New("invalid field name; only numbers, hyphens and lowercase letters are allowed")
var ErrInvalidLengthField = errors.New("field length is higher than the maximum")
var ErrInvalidWorkflowType = errors.New("invalid workflow type, must be either 'data', 'training' 'feedback' or 'serving'")
var ErrInvalidProcessType = errors.New("invalid process type, must be either 'trigger', 'task' or 'exit'")

func MissingRequiredFieldError(field string) error {
	return fmt.Errorf("%w: %s", ErrMissingRequiredField, field)
}

func InvalidFieldNameError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidFieldName, field)
}

func InvalidLengthFieldError(field string, maxLength int) error {
	return fmt.Errorf("%w: %s; maximum allowed: %d", ErrInvalidLengthField, field, maxLength)
}

func InvalidWorkflowTypeError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidWorkflowType, field)
}

func InvalidProcessTypeError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidProcessType, field)
}
