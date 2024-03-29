//go:build unit

package parse_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
	"github.com/konstellation-io/krt/pkg/parse"
)

func TestCorrectKrtFile(t *testing.T) {
	krt, err := parse.ParseFileToKrt("./testdata/correct_krt.yaml")
	require.NoError(t, err)

	err = krt.Validate()
	require.NoError(t, err)
}

func TestNonExistentFile(t *testing.T) {
	krt, err := parse.ParseFileToKrt("./testdata/non_existent_krt.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrReadingFile)
	assert.Nil(t, krt)
}

func TestInvalidFile(t *testing.T) {
	krt, err := parse.ParseFileToKrt("./testdata/invalid_file.yaml")
	require.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrInvalidYaml)
	assert.Nil(t, krt)
}

func TestCorrectKrtFileSettingDefaults(t *testing.T) {
	parsedKrt, err := parse.ParseFileToKrt("./testdata/missing_defaults_krt.yaml")
	require.NoError(t, err)

	err = parsedKrt.Validate()
	require.NoError(t, err)

	for idxWorkflow, workflows := range parsedKrt.Workflows {
		for idxProcess, process := range workflows.Processes {
			if idxWorkflow == 0 && idxProcess == 0 {
				assert.True(t, *process.GPU)
				assert.Equal(t, 2, *process.Replicas)
				assert.Equal(t, process.Networking.Protocol, krt.DefaultProtocol)
				assert.Equal(t, process.ResourceLimits.CPU.Request, process.ResourceLimits.CPU.Limit)
				assert.Equal(t, process.ResourceLimits.Memory.Request, process.ResourceLimits.Memory.Limit)
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
	parsedKrt, err := parse.ParseFileToKrt("./testdata/not_valid_krt.yaml")
	require.NoError(t, err)

	err = parsedKrt.Validate()
	require.Error(t, err)

	errList := strings.Split(err.Error(), "\n")
	require.Len(t, errList, 12)

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

	assert.ErrorIs(t, err, errors.ErrMissingRequiredField)
	assert.Contains(t, err.Error(), "krt.workflows[0].processes[0].resourceLimits")
}

func TestNotValidTypesKrt(t *testing.T) {
	parsedKrt, err := parse.ParseFileToKrt("./testdata/not_valid_types_krt.yaml")
	assert.NoError(t, err)

	err = parsedKrt.Validate()
	assert.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrInvalidProcessType)
	assert.ErrorIs(t, err, errors.ErrInvalidWorkflowType)
	assert.ErrorIs(t, err, errors.ErrInvalidProcessObjectStoreScope)
	assert.ErrorIs(t, err, errors.ErrInvalidNetworkingProtocol)
}

func TestValidKrtToYaml(t *testing.T) {
	krtYml, err := os.ReadFile("./testdata/correct_krt.yaml")
	require.NoError(t, err)

	krt, err := parse.ParseYamlToKrt(krtYml)
	require.NoError(t, err)

	yaml, err := parse.ParseKrtToYaml(krt)
	require.NoError(t, err)

	expectedYamlString := string(krtYml)
	expectedYamlString = strings.ReplaceAll(expectedYamlString, "\n", "")
	expectedYamlString = strings.ReplaceAll(expectedYamlString, " ", "")

	actualYamlString := string(yaml)
	actualYamlString = strings.ReplaceAll(actualYamlString, "\n", "")
	actualYamlString = strings.ReplaceAll(actualYamlString, " ", "")

	assert.Equal(t, expectedYamlString, actualYamlString)
}
