//go:build unit

package main

import (
	errorUtils "errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konstellation-io/krt/errors"
)

type test struct {
	name        string
	krtYaml     *Krt
	wantError   bool
	errorType   error
	errorString string
}

func TestKrtValidator(t *testing.T) {
	correctBuildTests := []test{
		{
			name:      "KRT YAML values successfully validated",
			krtYaml:   NewKrtBuilder().Build(),
			wantError: false,
		},
	}

	requiredFieldsTests := []test{
		{
			name:        "fails if krt hasn't required field name",
			krtYaml:     NewKrtBuilder().WithName("").Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.name").Error(),
		},
		{
			name:        "fails if krt hasn't required field description",
			krtYaml:     NewKrtBuilder().WithDescription("").Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.description").Error(),
		},
		{
			name:        "fails if krt hasn't required field version",
			krtYaml:     NewKrtBuilder().WithVersion("").Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.version").Error(),
		},
		{
			name:        "fails if krt hasn't required workflows declared",
			krtYaml:     NewKrtBuilder().WithWorkflows(nil).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows").Error(),
		},
		{
			name:        "fails if krt hasn't required workflow name",
			krtYaml:     NewKrtBuilder().WithWorkflowName("").Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].name").Error(),
		},
		{
			name:        "fails if krt hasn't required processes declared in a workflow",
			krtYaml:     NewKrtBuilder().WithProcesses(nil).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes").Error(),
		},
		{
			name:        "fails if krt hasn't required process name",
			krtYaml:     NewKrtBuilder().WithProcessName("", 0).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].name").Error(),
		},
		{
			name:        "fails if krt hasn't required process subscriptions",
			krtYaml:     NewKrtBuilder().WithProcessSubscriptions(nil, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].subscriptions").Error(),
		},
	}

	invalidNameTests := []test{
		{
			name:        "fails if version name has an invalid format",
			krtYaml:     NewKrtBuilder().WithVersion("Invalid string!").Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidFieldName,
			errorString: errors.InvalidFieldNameError("krt.version").Error(),
		},
		{
			name:        "fails if version name has an invalid length",
			krtYaml:     NewKrtBuilder().WithVersion("this-version-name-length-is-higher-than-the-maximum").Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.version", maxFieldNameLength).Error(),
		},
		{
			name:        "fails if krt workflow name has an invalid format",
			krtYaml:     NewKrtBuilder().WithWorkflowName("Invalid string!").Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidFieldName,
			errorString: errors.InvalidFieldNameError("krt.workflows[0].name").Error(),
		},
		{
			name: "fails if krt workflow name has an invalid length",
			krtYaml: NewKrtBuilder().WithWorkflowName(
				"this-workflow-name-length-is-higher-than-the-maximum",
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].name", maxFieldNameLength).Error(),
		},
		{
			name:        "fails if krt process name has an invalid format",
			krtYaml:     NewKrtBuilder().WithProcessName("Invalid string!", 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidFieldName,
			errorString: errors.InvalidFieldNameError("krt.workflows[0].processes[0].name").Error(),
		},
		{
			name: "fails if krt process name has an invalid length",
			krtYaml: NewKrtBuilder().WithProcessName(
				"this-process-name-length-is-higher-than-the-maximum",
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].processes[0].name", maxFieldNameLength).Error(),
		},
	}

	invalidTypeTests := []test{
		{
			name:        "fails if krt hasn't a valid workflow type",
			krtYaml:     NewKrtBuilder().WithWorkflowType("").Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidWorkflowType,
			errorString: errors.InvalidWorkflowTypeError("krt.workflows[0].type").Error(),
		},
		{
			name:        "fails if krt hasn't a valid process type",
			krtYaml:     NewKrtBuilder().WithProcessType("", 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessType,
			errorString: errors.InvalidProcessTypeError("krt.workflows[0].processes[0].type").Error(),
		},
	}

	allTests := make([]test, 0)
	allTests = append(allTests, correctBuildTests...)
	allTests = append(allTests, requiredFieldsTests...)
	allTests = append(allTests, invalidNameTests...)
	allTests = append(allTests, invalidTypeTests...)

	for _, tc := range allTests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.krtYaml.Validate()
			if tc.wantError {
				assert.True(t, errorUtils.Is(err, tc.errorType))
				assert.Equal(t, tc.errorString, err.Error())
			} else {
				assert.Empty(t, err)
			}
		})
	}
}
