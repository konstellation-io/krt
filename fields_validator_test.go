//go:build unit

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlFieldsValidator_Run(t *testing.T) {
	tests := []struct {
		name        string
		krtYaml     *Krt
		wantError   bool
		errorString string
	}{
		{
			name:        "KRT YAML values successfully validated",
			krtYaml:     NewKrtBuilder().Build(),
			wantError:   false,
			errorString: "",
		},
		{
			name:        "fails if krt hasn't required field name",
			krtYaml:     NewKrtBuilder().WithName("").Build(),
			wantError:   true,
			errorString: "the field \"Name\" is required",
		},
		{
			name:        "fails if krt hasn't required field description",
			krtYaml:     NewKrtBuilder().WithDescription("").Build(),
			wantError:   true,
			errorString: "the field \"Description\" is required",
		},
		{
			name:        "fails if krt hasn't required field version",
			krtYaml:     NewKrtBuilder().WithVersion("").Build(),
			wantError:   true,
			errorString: "the field \"Version\" is required",
		},
		{
			name:        "fails if version name has an invalid length",
			krtYaml:     NewKrtBuilder().WithVersion("this-version-name-length-is-higher-than-the-maximum").Build(),
			wantError:   true,
			errorString: "invalid length \"this-version-name-length-is-higher-than-the-maximum\" at \"Version\" must be lower than 20",
		},
	}

	valuesValidator := NewYamlFieldsValidator()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := valuesValidator.Run(tc.krtYaml)
			if tc.wantError {
				assert.Error(t, err[0], tc.errorString)
				return
			}

			assert.Empty(t, err)
		})
	}
}
