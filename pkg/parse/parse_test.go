//go:build unit

package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
)

func TestCorrectKrtFile(t *testing.T) {
	krt, err := ParseFile("./test_files/correct_krt.yaml")
	assert.NoError(t, err)

	err = krt.Validate()
	assert.NoError(t, err)
}

func TestCorrectKrtFileSettingDefaults(t *testing.T) {
	parsedKrt, err := ParseFile("./test_files/missing_defaults_krt.yaml")
	assert.NoError(t, err)

	err = parsedKrt.Validate()
	assert.NoError(t, err)

	for idxWorkflow, workflows := range parsedKrt.Workflows {
		for idxProcess, process := range workflows.Processes {
			if idxWorkflow == 0 && idxProcess == 0 {
				assert.True(t, *process.GPU)
				assert.Equal(t, 2, *process.Replicas)
				assert.Equal(t, krt.NetworkingProtocolUDP, process.Networking.DestinationProtocol)
				assert.Equal(t, krt.NetworkingProtocolUDP, process.Networking.TargetProtocol)
			} else if idxWorkflow == 0 && idxProcess == 1 {
				assert.Nil(t, process.Networking)
			} else {
				assert.NotNil(t, process.GPU)
				require.NotNil(t, process.Replicas)
				assert.GreaterOrEqual(t, *process.Replicas, 1)
				if process.Networking != nil {
					assert.NotEmpty(t, process.Networking.TargetPort)
					assert.NotEmpty(t, process.Networking.TargetProtocol)
					assert.NotEmpty(t, process.Networking.DestinationPort)
					assert.NotEmpty(t, process.Networking.DestinationProtocol)
				}
			}
		}
	}
}

func TestNotValidTypesKrt(t *testing.T) {
	parsedKrt, err := ParseFile("./test_files/not_valid_types_krt.yaml")
	assert.NoError(t, err)

	err = parsedKrt.Validate()
	assert.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrInvalidProcessType)
	assert.ErrorIs(t, err, errors.ErrInvalidWorkflowType)
	assert.ErrorIs(t, err, errors.ErrInvalidProcessObjectStoreScope)
	assert.ErrorIs(t, err, errors.ErrInvalidNetworkingProtocol)
}
