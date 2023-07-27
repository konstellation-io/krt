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
var ErrInvalidProcessCPUResourceLimit = errors.New("invalid process CPU resource limit, must be of form '1', '0.5' or '100m'")
var ErrInvalidProcessCPURelation = errors.New("invalid process CPU, 'limit' cannot be lower than 'request'")
var ErrInvalidProcessMemoryResourceLimit = errors.New("invalid process memory resource limit, must be of form '350M' or '1Gi'")
var ErrInvalidProcessMemoryRelation = errors.New("invalid process memory, 'limit' cannot be lower than 'request'")

var ErrNotEnoughProcesses = errors.New("not enough processes declared for this workflow, needed at least 1 trigger and 1 exit process")
var ErrDuplicatedProcessName = errors.New("process names must be unique")
var ErrDuplicatedProcessSubscription = errors.New("subscriptions cannot be duplicated")
var ErrInvalidProcessSubscription = errors.New("invalid subscription")
var ErrCannotSubscribeToItself = errors.New("cannot subscribe to itself")
var ErrCannotSubscribeToNonExistentProcess = errors.New("cannot subscribe to non existent process")

func errorWithMessage(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}

func MissingRequiredFieldError(field string) error {
	return errorWithMessage(ErrMissingRequiredField, field)
}

func InvalidVersionTagError(field string) error {
	return errorWithMessage(ErrInvalidVersionTag, field)
}

func InvalidFieldNameError(field string) error {
	return errorWithMessage(ErrInvalidFieldName, field)
}

func InvalidLengthFieldError(field string, maxLength int) error {
	return fmt.Errorf("%w: %s; maximum length allowed: %d", ErrInvalidLengthField, field, maxLength)
}

func DuplicatedWorkflowNameError(field string) error {
	return errorWithMessage(ErrDuplicatedWorkflowName, field)
}

func InvalidWorkflowTypeError(field string) error {
	return errorWithMessage(ErrInvalidWorkflowType, field)
}

func InvalidProcessTypeError(field string) error {
	return errorWithMessage(ErrInvalidProcessType, field)
}

func InvalidProcessObjectStoreScopeError(field string) error {
	return errorWithMessage(ErrInvalidProcessObjectStoreScope, field)
}

func InvalidNetworkingProtocolError(field string) error {
	return errorWithMessage(ErrInvalidNetworkingProtocol, field)
}

func InvalidProcessCPUError(field string) error {
	return errorWithMessage(ErrInvalidProcessCPU, field)
}

func InvalidProcessCPURelationError(field string) error {
	return errorWithMessage(ErrInvalidProcessCPURelation, field)
}

func InvalidProcessMemoryError(field string) error {
	return errorWithMessage(ErrInvalidProcessMemory, field)
}

func InvalidProcessMemoryRelationError(field string) error {
	return errorWithMessage(ErrInvalidProcessMemoryRelation, field)
}

func NotEnoughProcessesError(field string) error {
	return errorWithMessage(ErrNotEnoughProcesses, field)
}

func DuplicatedProcessNameError(field string) error {
	return errorWithMessage(ErrDuplicatedProcessName, field)
}

func DuplicatedProcessSubscriptionError(field string) error {
	return errorWithMessage(ErrDuplicatedProcessSubscription, field)
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
