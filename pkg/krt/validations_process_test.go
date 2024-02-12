//go:build unit

package krt_test

import (
	"testing"

	"github.com/konstellation-io/krt/internal/kubeutil"
	"github.com/konstellation-io/krt/pkg/krt"
	"github.com/stretchr/testify/assert"
)

func TestProcess_Validate(t *testing.T) {
	testCases := []struct {
		name          string
		process       krt.Process
		expectedError error
	}{
		{
			"valid process with node selectors",
			*NewProcessBuilder().WithNodeSelectors(map[string]string{"valid-key": "valid-value"}).Build(),
			nil,
		},
		{
			"process with invalid node selector key prefix",
			*NewProcessBuilder().WithNodeSelectors(map[string]string{"inalid prefix/key-name": "valid-value"}).Build(),
			kubeutil.ErrInvalidKeyPrefix,
		},
		{
			"process with valid node selector key value",
			*NewProcessBuilder().WithNodeSelectors(map[string]string{"key-name": "valid-value"}).Build(),
			nil,
		},
		{
			"process with invalid node selector key value",
			*NewProcessBuilder().WithNodeSelectors(map[string]string{"key-name": "invalid value"}).Build(),
			kubeutil.ErrInvalidValue,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ErrorIs(t, tc.process.Validate(0, 0), tc.expectedError)
		})
	}
}
