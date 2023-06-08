//go:build unit

package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errTest = errors.New("test error")
var errTestAlternative = errors.New("test error alternative")

func wrapError(err error) error {
	return fmt.Errorf("%w: wrap error", err)
}

func TestErrorsJoin(t *testing.T) {
	err := Join(
		errTest,
		errTestAlternative,
	)

	assert.ErrorIs(t, err, errTest)
	assert.ErrorIs(t, err, errTestAlternative)
}

func TestIs(t *testing.T) {
	err := wrapError(errTest)

	assert.Equal(t, errors.Is(err, errTest), Is(err, errTest))
}

func TestMissingRequiredFieldError(t *testing.T) {
	err := MissingRequiredFieldError("test")

	assert.ErrorIs(t, err, ErrMissingRequiredField)
}

func TestInvalidFieldNameError(t *testing.T) {
	err := InvalidFieldNameError("test")

	assert.ErrorIs(t, err, ErrInvalidFieldName)
}

func TestInvalidLengthFieldError(t *testing.T) {
	err := InvalidLengthFieldError("test", 10)

	assert.ErrorIs(t, err, ErrInvalidLengthField)
}

func TestDuplicatedWorkflowNameError(t *testing.T) {
	err := DuplicatedWorkflowNameError("test")

	assert.ErrorIs(t, err, ErrDuplicatedWorkflowName)
}

func TestInvalidWorkflowTypeError(t *testing.T) {
	err := InvalidWorkflowTypeError("test")

	assert.ErrorIs(t, err, ErrInvalidWorkflowType)
}

func TestInvalidProcessTypeError(t *testing.T) {
	err := InvalidProcessTypeError("test")

	assert.ErrorIs(t, err, ErrInvalidProcessType)
}

func TestInvalidProcessObjectStoreScopeError(t *testing.T) {
	err := InvalidProcessObjectStoreScopeError("test")

	assert.ErrorIs(t, err, ErrInvalidProcessObjectStoreScope)
}

func TestInvalidNetworkingProtocolError(t *testing.T) {
	err := InvalidNetworkingProtocolError("test")

	assert.ErrorIs(t, err, ErrInvalidNetworkingProtocol)
}

func TestNotEnoughProcessesError(t *testing.T) {
	err := NotEnoughProcessesError("test")

	assert.ErrorIs(t, err, ErrNotEnoughProcesses)
}

func TestDuplicatedProcessNameError(t *testing.T) {
	err := DuplicatedProcessNameError("test")

	assert.ErrorIs(t, err, ErrDuplicatedProcessName)
}

func TestDuplicatedProcessSubscriptionError(t *testing.T) {
	err := DuplicatedProcessSubscriptionError("test")

	assert.ErrorIs(t, err, ErrDuplicatedProcessSubscription)
}

func TestInvalidProcessSubscriptionError(t *testing.T) {
	err := InvalidProcessSubscriptionError("test", "test", "test")

	assert.ErrorIs(t, err, ErrInvalidProcessSubscription)
}

func TestCannotSubscribeToItselfError(t *testing.T) {
	err := CannotSubscribeToItselfError("test")

	assert.ErrorIs(t, err, ErrCannotSubscribeToItself)
}

func TestInvalidYamlError(t *testing.T) {
	err := InvalidYamlError(errTest)

	assert.ErrorIs(t, err, ErrInvalidYaml)
}

func TestReadingFileError(t *testing.T) {
	err := ReadingFileError(errTest)

	assert.ErrorIs(t, err, ErrReadingFile)
}
