//go:build unit

package errors_test

import (
	errUtils "errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konstellation-io/krt/pkg/errors"
)

var errTest = errUtils.New("test error")
var errTestAlternative = errUtils.New("test error alternative")

func wrapError(err error) error {
	return fmt.Errorf("%w: wrap error", err)
}

func TestErrorsJoin(t *testing.T) {
	err := errors.Join(
		errTest,
		errTestAlternative,
	)

	assert.ErrorIs(t, err, errTest)
	assert.ErrorIs(t, err, errTestAlternative)
}

func TestIs(t *testing.T) {
	err := wrapError(errTest)
	assert.Equal(t, errUtils.Is(err, errTest), errors.Is(err, errTest))
}

func TestMissingRequiredFieldError(t *testing.T) {
	err := errors.MissingRequiredFieldError("test")
	assert.ErrorIs(t, err, errors.ErrMissingRequiredField)
}

func TestInvalidVersionTagError(t *testing.T) {
	err := errors.InvalidVersionTagError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidVersionTag)
}

func TestInvalidFieldNameError(t *testing.T) {
	err := errors.InvalidFieldNameError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidFieldName)
}

func TestInvalidLengthFieldError(t *testing.T) {
	err := errors.InvalidLengthFieldError("test", 10)
	assert.ErrorIs(t, err, errors.ErrInvalidLengthField)
}

func TestDuplicatedWorkflowNameError(t *testing.T) {
	err := errors.DuplicatedWorkflowNameError("test")
	assert.ErrorIs(t, err, errors.ErrDuplicatedWorkflowName)
}

func TestInvalidWorkflowTypeError(t *testing.T) {
	err := errors.InvalidWorkflowTypeError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidWorkflowType)
}

func TestInvalidProcessTypeError(t *testing.T) {
	err := errors.InvalidProcessTypeError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidProcessType)
}

func TestInvalidProcessObjectStoreScopeError(t *testing.T) {
	err := errors.InvalidProcessObjectStoreScopeError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidProcessObjectStoreScope)
}

func TestInvalidNetworkingProtocolError(t *testing.T) {
	err := errors.InvalidNetworkingProtocolError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidNetworkingProtocol)
}

func TestInvalidProcessCPUError(t *testing.T) {
	err := errors.InvalidProcessCPUError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidProcessCPU)
}

func TestInvalidProcessCPURelationError(t *testing.T) {
	err := errors.InvalidProcessCPURelationError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidProcessCPURelation)
}

func TestInvalidProcessMemoryError(t *testing.T) {
	err := errors.InvalidProcessMemoryError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidProcessMemory)
}

func TestInvalidProcessMemoryRelationError(t *testing.T) {
	err := errors.InvalidProcessMemoryRelationError("test")
	assert.ErrorIs(t, err, errors.ErrInvalidProcessMemoryRelation)
}

func TestNotEnoughProcessesError(t *testing.T) {
	err := errors.NotEnoughProcessesError("test")
	assert.ErrorIs(t, err, errors.ErrNotEnoughProcesses)
}

func TestDuplicatedProcessNameError(t *testing.T) {
	err := errors.DuplicatedProcessNameError("test")
	assert.ErrorIs(t, err, errors.ErrDuplicatedProcessName)
}

func TestDuplicatedProcessSubscriptionError(t *testing.T) {
	err := errors.DuplicatedProcessSubscriptionError("test")
	assert.ErrorIs(t, err, errors.ErrDuplicatedProcessSubscription)
}

func TestInvalidProcessSubscriptionError(t *testing.T) {
	err := errors.InvalidProcessSubscriptionError("test", "test", "test")
	assert.ErrorIs(t, err, errors.ErrInvalidProcessSubscription)
}

func TestCannotSubscribeToItselfError(t *testing.T) {
	err := errors.CannotSubscribeToItselfError("test")
	assert.ErrorIs(t, err, errors.ErrCannotSubscribeToItself)
}

func TestCannotSubscribeToNonExistentProcessError(t *testing.T) {
	err := errors.CannotSubscribeToNonExistentProcessError("test-process", "test")
	assert.ErrorIs(t, err, errors.ErrCannotSubscribeToNonExistentProcess)
}

func TestInvalidYamlError(t *testing.T) {
	err := errors.InvalidYamlError(errTest)
	assert.ErrorIs(t, err, errors.ErrInvalidYaml)
}

func TestReadingFileError(t *testing.T) {
	err := errors.ReadingFileError(errTest)
	assert.ErrorIs(t, err, errors.ErrReadingFile)
}
