package errors

import (
	"errors"
	"fmt"
	"strings"
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

func IsErrorStringInError(expectedErrorString string, err error) bool {
	allErrorStrings := strings.Split(err.Error(), "\n")
	for _, errorString := range allErrorStrings {
		if errorString == expectedErrorString {
			return true
		}
	}
	return false
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

var ErrMissingRequiredField = errors.New("missing required field")
var ErrInvalidFieldName = errors.New("invalid field name; only numbers, hyphens and lowercase letters are allowed")
var ErrInvalidLengthField = errors.New("field length is higher than the maximum")
var ErrInvalidWorkflowType = errors.New("invalid workflow type, must be either 'data', 'training' 'feedback' or 'serving'")
var ErrInvalidProcessType = errors.New("invalid process type, must be either 'trigger', 'task' or 'exit'")
var ErrInvalidProcessBuild = errors.New("invalid process build, must have either 'image' or 'dockerfile'")
var ErrInvalidProcessObjectStoreScope = errors.New("invalid process object store scope, must be either 'product' or 'workflow'")
var ErrInvalidNetworkingProtocol = errors.New("invalid networking protocol, must be either 'UDP' or 'TCP'")
var ErrDuplicatedProcessSubscription = errors.New("subscriptions cannot be duplicated")
var ErrInvalidProcessSubscription = errors.New("invalid subscription")

func MissingRequiredFieldError(field string) error {
	return fmt.Errorf("%w: %s", ErrMissingRequiredField, field)
}

func InvalidFieldNameError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidFieldName, field)
}

func InvalidLengthFieldError(field string, maxLength int) error {
	return fmt.Errorf("%w: %s; maximum length allowed: %d", ErrInvalidLengthField, field, maxLength)
}

func InvalidWorkflowTypeError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidWorkflowType, field)
}

func InvalidProcessTypeError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidProcessType, field)
}

func InvalidProcessBuildError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidProcessBuild, field)
}

func InvalidProcessObjectStoreScopeError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidProcessObjectStoreScope, field)
}

func InvalidNetworkingProtocolError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidNetworkingProtocol, field)
}

func DuplicatedProcessSubscriptionError(field string) error {
	return fmt.Errorf("%w: %s", ErrDuplicatedProcessSubscription, field)
}

func InvalidProcessSubscriptionError(processType, subscritpionProcessType, field string) error {
	return fmt.Errorf("%w: this process of type %q cannot subscribe to %q processes, in %s", ErrInvalidProcessSubscription, processType, subscritpionProcessType, field)
}
