package errors

import (
	"errors"
	"fmt"
)

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

// Validation errors.

var ErrMissingRequiredField = errors.New("missing required field")
var ErrInvalidVersionTag = errors.New("invalid version tag; must follow the format 'vX.Y.Z'")
var ErrInvalidFieldName = errors.New("invalid field name; only numbers, hyphens and lowercase letters are allowed")
var ErrInvalidLengthField = errors.New("field length is higher than the maximum")

var ErrDuplicatedWorkflowName = errors.New("workflow names must be unique")
var ErrInvalidWorkflowType = errors.New("invalid workflow type, must be either 'data', 'training' 'feedback' or 'serving'")

var ErrInvalidProcessType = errors.New("invalid process type, must be either 'trigger', 'task' or 'exit'")
var ErrInvalidProcessObjectStoreScope = errors.New("invalid process object store scope, must be either 'product' or 'workflow'")
var ErrInvalidNetworkingProtocol = errors.New("invalid networking protocol, must be either 'UDP' or 'TCP'")

var ErrNotEnoughProcesses = errors.New("not enough processes declared for this workflow, needed at least 1 trigger and 1 exit process")
var ErrDuplicatedProcessName = errors.New("process names must be unique")
var ErrDuplicatedProcessSubscription = errors.New("subscriptions cannot be duplicated")
var ErrInvalidProcessSubscription = errors.New("invalid subscription")
var ErrCannotSubscribeToItself = errors.New("cannot subscribe to itself")
var ErrCannotSubscribeToNonExistentProcess = errors.New("cannot subscribe to non existent process")

func MissingRequiredFieldError(field string) error {
	return fmt.Errorf("%w: %s", ErrMissingRequiredField, field)
}

func InvalidVersionTagError(field string) error {
	return fmt.Errorf("%w: %s", ErrInvalidVersionTag, field)
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
	return fmt.Errorf(
		"%w: this process of type %q cannot subscribe to %q processes, in %s",
		ErrInvalidProcessSubscription,
		processType,
		subscritpionProcessType,
		field,
	)
}

func CannotSubscribeToItselfError(field string) error {
	return fmt.Errorf("%w: %s", ErrCannotSubscribeToItself, field)
}

func CannotSubscribeToNonExistentProcessError(process, field string) error {
	return fmt.Errorf("%w: process named %q does not exist %s", ErrCannotSubscribeToNonExistentProcess, process, field)
}

// Parse errors.

var ErrInvalidYaml = errors.New("invalid yaml")
var ErrReadingFile = errors.New("error reading file")

func InvalidYamlError(err error) error {
	return fmt.Errorf("error unmarshalling krt yaml, %w: %w", ErrInvalidYaml, err)
}

func ReadingFileError(err error) error {
	return fmt.Errorf("%w: %w", ErrReadingFile, err)
}
