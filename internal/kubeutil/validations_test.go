package kubeutil_test

import (
	"testing"

	"github.com/konstellation-io/krt/internal/kubeutil"
	"github.com/stretchr/testify/assert"
)

func TestValidateNodeSelectorKey(t *testing.T) {
	testCases := []struct {
		name          string
		key           string
		expectedError error
	}{
		{"Valid key without prefix", "valid-key", nil},
		{"Valid key with prefix", "konstellation.io/valid-key", nil},
		{"Invalid key with empty prefix", "/valid-key", kubeutil.ErrInvalidKeyPrefix},
		{"Invalid key without prefix", "invalid key", kubeutil.ErrInvalidKeyName},
		{"Invalid key with prefix", "invalid prefix/invalid key", kubeutil.ErrInvalidKeyPrefix},
		{"Invalid key format multiple '/'", "invalid/key/format", kubeutil.ErrInvalidKeyFormat},
		{"Invalid empty key", "", kubeutil.ErrInvalidKeyName},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ErrorIs(t, kubeutil.ValidateNodeSelectorKey(tc.key), tc.expectedError)
		})
	}
}

func TestValidateNodeSelectorValue(t *testing.T) {
	testCases := []struct {
		name          string
		value         string
		expectedError error
	}{
		{"Valid value", "valid-value", nil},
		{"Invalid value", "invalid value", kubeutil.ErrInvalidValue},
		{"Invalid empty value", "", kubeutil.ErrInvalidValue},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ErrorIs(t, kubeutil.ValidateNodeSelectorValue(tc.value), tc.expectedError)
		})
	}
}
