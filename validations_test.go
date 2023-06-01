//go:build unit

package main

import (
	errorUtils "errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/krt/errors"
)

func TestKrtValidator(t *testing.T) {
	tests := []struct {
		name      string
		krtYaml   *Krt
		wantError bool
		errorType error
	}{
		// Correct Build
		{
			name:      "KRT YAML values successfully validated",
			krtYaml:   NewKrtBuilder().Build(),
			wantError: false,
		},
		// Version related
		{
			name:      "fails if krt hasn't required field name",
			krtYaml:   NewKrtBuilder().WithName("").Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
		{
			name:      "fails if krt hasn't required field description",
			krtYaml:   NewKrtBuilder().WithDescription("").Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
		{
			name:      "fails if krt hasn't required field version",
			krtYaml:   NewKrtBuilder().WithVersion("").Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
		{
			name:      "fails if version name has an invalid format",
			krtYaml:   NewKrtBuilder().WithVersion("Invalid string!").Build(),
			wantError: true,
			errorType: errors.ErrInvalidFieldName,
		},
		{
			name:      "fails if version name has an invalid length",
			krtYaml:   NewKrtBuilder().WithVersion("this-version-name-length-is-higher-than-the-maximum").Build(),
			wantError: true,
			errorType: errors.ErrInvalidLengthField,
		},
		{
			name:      "fails if krt hasn't required workflows declared",
			krtYaml:   NewKrtBuilder().WithWorkflows(nil).Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
		// Workflow related
		{
			name:      "fails if krt hasn't required workflow name",
			krtYaml:   NewKrtBuilder().WithWorkflowName("").Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
		{
			name:      "fails if krt workflow name has an invalid format",
			krtYaml:   NewKrtBuilder().WithWorkflowName("Invalid string!").Build(),
			wantError: true,
			errorType: errors.ErrInvalidFieldName,
		},
		{
			name: "fails if krt workflow name has an invalid length",
			krtYaml: NewKrtBuilder().WithWorkflowName(
				"this-workflow-name-length-is-higher-than-the-maximum",
			).Build(),
			wantError: true,
			errorType: errors.ErrInvalidLengthField,
		},
		{
			name:      "fails if krt hasn't a valid workflow type",
			krtYaml:   NewKrtBuilder().WithWorkflowType("").Build(),
			wantError: true,
			errorType: errors.ErrInvalidWorkflowType,
		},
		{
			name:      "fails if krt hasn't required processes declared in a workflow",
			krtYaml:   NewKrtBuilder().WithProcesses(nil).Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
		// Process related
		{
			name:      "fails if krt hasn't required process name",
			krtYaml:   NewKrtBuilder().WithProcessName("", 0).Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
		{
			name:      "fails if krt process name has an invalid format",
			krtYaml:   NewKrtBuilder().WithProcessName("Invalid string!", 0).Build(),
			wantError: true,
			errorType: errors.ErrInvalidFieldName,
		},
		{
			name: "fails if krt process name has an invalid length",
			krtYaml: NewKrtBuilder().WithProcessName(
				"this-process-name-length-is-higher-than-the-maximum",
				0,
			).Build(),
			wantError: true,
			errorType: errors.ErrInvalidLengthField,
		},
		{
			name:      "fails if krt hasn't a valid process type",
			krtYaml:   NewKrtBuilder().WithProcessType("", 0).Build(),
			wantError: true,
			errorType: errors.ErrInvalidProcessType,
		},
		{
			name:      "fails if krt hasn't required process subscriptions",
			krtYaml:   NewKrtBuilder().WithProcessSubscriptions(nil, 0).Build(),
			wantError: true,
			errorType: errors.ErrMissingRequiredField,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errs := tc.krtYaml.Validate()
			if tc.wantError {
				require.Len(t, errs, 1)
				assert.Error(t, errs[0])
				assert.True(t, errorUtils.Is(errs[0], tc.errorType))
			} else {
				assert.Empty(t, errs)
			}
		})
	}
}
