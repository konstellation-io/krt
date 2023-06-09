//go:build unit

package parse_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
	"github.com/konstellation-io/krt/pkg/parse"
)

func TestCorrectKrtFile(t *testing.T) {
	krt, err := parse.ParseFile("./test_files/correct_krt.yaml")
	require.NoError(t, err)

	err = krt.Validate()
	require.NoError(t, err)
}

func TestNonExistentFile(t *testing.T) {
	krt, err := parse.ParseFile("./test_files/non_existent_krt.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrReadingFile)
	assert.Nil(t, krt)
}

func TestInvalidFile(t *testing.T) {
	krt, err := parse.ParseFile("./test_files/invalid_file.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrInvalidYaml)
	assert.Nil(t, krt)
}

func TestCorrectKrtFileSettingDefaults(t *testing.T) {
	parsedKrt, err := parse.ParseFile("./test_files/missing_defaults_krt.yaml")
	require.NoError(t, err)

	err = parsedKrt.Validate()
	require.NoError(t, err)

	for idxWorkflow, workflows := range parsedKrt.Workflows {
		for idxProcess, process := range workflows.Processes {
			if idxWorkflow == 0 && idxProcess == 0 {
				assert.True(t, *process.GPU)
				assert.Equal(t, 2, *process.Replicas)
				assert.Equal(t, krt.NetworkingProtocolTCP, krt.DefaultProtocol)
			} else if idxWorkflow == 0 && idxProcess == 1 {
				assert.Nil(t, process.Networking)
			} else {
				assert.NotNil(t, process.GPU)
				assert.Equal(t, krt.DefaultGPUValue, *process.GPU)
				require.NotNil(t, process.Replicas)
				assert.Equal(t, krt.DefaultNumberOfReplicas, *process.Replicas)
				if process.Networking != nil {
					assert.NotEmpty(t, process.Networking.TargetPort)
					assert.NotEmpty(t, process.Networking.DestinationPort)
					assert.NotEmpty(t, process.Networking.Protocol)
				}
			}
		}
	}
}

func TestNotValidKrt(t *testing.T) {
	parsedKrt, err := parse.ParseFile("./test_files/not_valid_krt.yaml")
	require.NoError(t, err)

	err = parsedKrt.Validate()
	require.Error(t, err)

	errList := strings.Split(err.Error(), "\n")
	require.Len(t, errList, 10)

	assert.ErrorIs(t, err, errors.ErrInvalidVersionTag)
	assert.Contains(t, err.Error(), "krt.version")

	assert.ErrorIs(t, err, errors.ErrInvalidFieldName)
	assert.Contains(t, err.Error(), "krt.workflows[0].name")

	assert.ErrorIs(t, err, errors.ErrInvalidWorkflowType)
	assert.Contains(t, err.Error(), "krt.workflows[0].type")

	assert.ErrorIs(t, err, errors.ErrMissingRequiredField)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes[0].image")

	assert.ErrorIs(t, err, errors.ErrNotEnoughProcesses)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes")

	assert.ErrorIs(t, err, errors.ErrCannotSubscribeToItself)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes[0].subscriptions.entrypoint")

	assert.ErrorIs(t, err, errors.ErrInvalidProcessSubscription)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes[0].subscriptions.entrypoint")

	assert.ErrorIs(t, err, errors.ErrInvalidFieldName)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes[1].name")

	assert.ErrorIs(t, err, errors.ErrInvalidProcessType)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes[1].type")

	assert.ErrorIs(t, err, errors.ErrCannotSubscribeToNonExistentProcess)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes[1]")

	assert.ErrorIs(t, err, errors.ErrNotEnoughProcesses)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes")
}

func TestNotValidTypesKrt(t *testing.T) {
	parsedKrt, err := parse.ParseFile("./test_files/not_valid_types_krt.yaml")
	assert.NoError(t, err)

	err = parsedKrt.Validate()
	assert.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrInvalidProcessType)
	assert.ErrorIs(t, err, errors.ErrInvalidWorkflowType)
	assert.ErrorIs(t, err, errors.ErrInvalidProcessObjectStoreScope)
	assert.ErrorIs(t, err, errors.ErrInvalidNetworkingProtocol)
}
