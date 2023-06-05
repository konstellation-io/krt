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

func Is(err, target error) bool {
	return errors.Is(err, target)
}

var ErrMissingRequiredField = errors.New("missing required field")
var ErrInvalidFieldName = errors.New("invalid field name; only numbers, hyphens and lowercase letters are allowed")
var ErrInvalidLengthField = errors.New("field length is higher than the maximum")

var ErrDuplicatedWorkflowName = errors.New("workflow names must be unique")
var ErrInvalidWorkflowType = errors.New("invalid workflow type, must be either 'data', 'training' 'feedback' or 'serving'")

var ErrInvalidProcessType = errors.New("invalid process type, must be either 'trigger', 'task' or 'exit'")
var ErrInvalidProcessBuild = errors.New("invalid process build, must have either 'image' or 'dockerfile'")
var ErrInvalidProcessObjectStoreScope = errors.New("invalid process object store scope, must be either 'product' or 'workflow'")
var ErrInvalidNetworkingProtocol = errors.New("invalid networking protocol, must be either 'UDP' or 'TCP'")

var ErrNotEnoughProcesses = errors.New("not enough processes declared for this workflow, needed at least 1 trigger and 1 exit process")
var ErrDuplicatedProcessName = errors.New("process names must be unique")
var ErrDuplicatedProcessSubscription = errors.New("subscriptions cannot be duplicated")
var ErrInvalidProcessSubscription = errors.New("invalid subscription")
var ErrCannotSubscribeToItself = errors.New("cannot subscribe to itself")

func MissingRequiredFieldError(field string) error {
	return fmt.Errorf("%w: %s", ErrMissingRequiredField, field)
}

func InvalidFieldNameError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidFieldName, field)
}

func InvalidLengthFieldError(field string, maxLength int) error {
	return fmt.Errorf("%w: %s; maximum length allowed: %d", ErrInvalidLengthField, field, maxLength)
}

func DuplicatedWorkflowNameError(field string) error {
	return fmt.Errorf("%w: %s", ErrDuplicatedWorkflowName, field)
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

func NotEnoughProcessesError(field string) error {
	return fmt.Errorf("%w: %s", ErrNotEnoughProcesses, field)
}

func DuplicatedProcessNameError(field string) error {
	return fmt.Errorf("%w: %s", ErrDuplicatedProcessName, field)
}

func DuplicatedProcessSubscriptionError(field string) error {
	return fmt.Errorf("%w: %s", ErrDuplicatedProcessSubscription, field)
}

func InvalidProcessSubscriptionError(processType, subscritpionProcessType, field string) error {
	return fmt.Errorf("%w: this process of type %q cannot subscribe to %q processes, in %s", ErrInvalidProcessSubscription, processType, subscritpionProcessType, field)
}

func CannotSubscribeToItselfError(field string) error {
	return fmt.Errorf("%w: %s", ErrCannotSubscribeToItself, field)
}
