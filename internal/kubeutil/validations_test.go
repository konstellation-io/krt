package kubeutil_test

import (
	"testing"

	"github.com/konstellation-io/krt/internal/kubeutil"
	"github.com/stretchr/testify/assert"
)

func TestValidateNodeSelectorKey(t *testing.T) {
	testCases := []struct {
		name      string
		key       string
		wantError bool
	}{
		{"Valid key without prefix", "valid-key", false},
		{"Valid key with prefix", "konstellation.io/valid-key", false},
		{"Invalid key without prefix", "invalid key", true},
		{"Invalid key with prefix", "invalid prefix/invalid key", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantError {
				assert.Error(t, kubeutil.ValidateNodeSelectorKey(tc.key))
			} else {
				assert.NoError(t, kubeutil.ValidateNodeSelectorKey(tc.key))
			}
		})
	}
}

func TestValidateNodeSelectorValue(t *testing.T) {
	testCases := []struct {
		name      string
		value     string
		wantError bool
	}{
		{"Valid value", "valid-value", false},
		{"Invalid value", "invalid value", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantError {
				assert.Error(t, kubeutil.ValidateNodeSelectorValue(tc.value))
			} else {
				assert.NoError(t, kubeutil.ValidateNodeSelectorValue(tc.value))
			}
		})
	}
}
