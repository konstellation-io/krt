//go:build unit

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestYamlFieldsValidator_Run(t *testing.T) {
	tests := []struct {
		name        string
		krtYaml     *Krt
		wantError   bool
		errorString string
	}{
		// Correct Build
		{
			name:        "KRT YAML values successfully validated",
			krtYaml:     NewKrtBuilder().Build(),
			wantError:   false,
			errorString: "",
		},
		// Version related
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
			name:        "fails if version name has an invalid format",
			krtYaml:     NewKrtBuilder().WithVersion("Invalid string!").Build(),
			wantError:   true,
			errorString: "invalid resource name \"Invalid string!\" at \"Version\"",
		},
		{
			name:        "fails if version name has an invalid length",
			krtYaml:     NewKrtBuilder().WithVersion("this-version-name-length-is-higher-than-the-maximum").Build(),
			wantError:   true,
			errorString: "invalid length \"this-version-name-length-is-higher-than-the-maximum\" at \"Version\" must be lower than 20",
		},
		{
			name:        "fails if krt hasn't required workflows declared",
			krtYaml:     NewKrtBuilder().WithWorkflows(nil).Build(),
			wantError:   true,
			errorString: "the field \"Workflows\" is required",
		},
		// Workflow related
		{
			name: "fails if krt hasn't required workflow name",
			krtYaml: NewKrtBuilder().WithWorkflows([]Workflow{
				{
					Name:      "",
					Type:      WorkflowTypeTraining,
					Processes: []Process{},
				},
			}).Build(),
			wantError:   true,
			errorString: "the field \"Workflows[0].Name\" is required",
		},
		{
			name: "fails if krt workflow name has an invalid format",
			krtYaml: NewKrtBuilder().WithWorkflows([]Workflow{
				{
					Name:      "Invalid string!",
					Type:      WorkflowTypeTraining,
					Processes: []Process{},
				},
			}).Build(),
			wantError:   true,
			errorString: "invalid resource name \"Invalid string!\" at \"Workflows[0].Name\"",
		},
		{
			name: "fails if krt workflow name has an invalid length",
			krtYaml: NewKrtBuilder().WithWorkflows([]Workflow{
				{
					Name:      "this-workflow-name-length-is-higher-than-the-maximum",
					Type:      WorkflowTypeTraining,
					Processes: []Process{},
				},
			}).Build(),
			wantError:   true,
			errorString: "invalid length \"this-workflow-name-length-is-higher-than-the-maximum\" at \"Workflows[0].Name\" must be lower than 20",
		},
		{
			name: "fails if krt hasn't required workflow type",
			krtYaml: NewKrtBuilder().WithWorkflows([]Workflow{
				{
					Name:      "test-workflow",
					Processes: []Process{},
				},
			}).Build(),
			wantError:   true,
			errorString: "the field \"Workflows[0].Type\" is required",
		},
		{
			name:        "fails if krt hasn't required processes declared in a workflow",
			krtYaml:     NewKrtBuilder().WithProcessesForWorkflow(nil, 0).Build(),
			wantError:   true,
			errorString: "the field \"Workflows[0].Processes\" is required",
		},
		// Process related
		{
			name: "fails if krt hasn't required process name",
			krtYaml: NewKrtBuilder().WithProcessesForWorkflow([]Process{
				{
					Name:          "",
					Type:          ProcessTypeTrigger,
					Build:         ProcessBuild{Image: "test-trigger-image"},
					Subscriptions: []string{},
				},
			}, 0).Build(),
			wantError:   true,
			errorString: "the field \"Workflows[0].Processes[0].Name\" is required",
		},
		{
			name: "fails if krt process name has an invalid format",
			krtYaml: NewKrtBuilder().WithProcessesForWorkflow([]Process{
				{
					Name:  "Invalid string!",
					Type:  ProcessTypeTrigger,
					Build: ProcessBuild{Image: "test-trigger-image"},
				},
			}, 0).Build(),
			wantError:   true,
			errorString: "invalid resource name \"Invalid string!\" at \"Workflows[0].Processes[0].Name\"",
		},
		{
			name: "fails if krt process name has an invalid length",
			krtYaml: NewKrtBuilder().WithProcessesForWorkflow([]Process{
				{
					Name:          "this-process-name-length-is-higher-than-the-maximum",
					Type:          ProcessTypeTrigger,
					Subscriptions: []string{},
				},
			}, 0).Build(),
			wantError:   true,
			errorString: "invalid length \"this-process-name-length-is-higher-than-the-maximum\" at \"Workflows[0].Processes[0].Name\" must be lower than 20",
		},
		{
			name: "fails if krt hasn't required process type",
			krtYaml: NewKrtBuilder().WithProcessesForWorkflow([]Process{
				{
					Name:          "test-process",
					Build:         ProcessBuild{Image: "test-trigger-image"},
					Subscriptions: []string{},
				},
			}, 0).Build(),
			wantError:   true,
			errorString: "the field \"Workflows[0].Processes[0].Type\" is required",
		},
		{
			name: "fails if krt hasn't required process subscriptions",
			krtYaml: NewKrtBuilder().WithProcessesForWorkflow([]Process{
				{
					Name:  "test-process",
					Type:  ProcessTypeTrigger,
					Build: ProcessBuild{Image: "test-trigger-image"},
				},
			}, 0).Build(),
			wantError:   true,
			errorString: "the field \"Workflows[0].Processes[0].Subscriptions\" is required",
		},
	}

	valuesValidator := NewYamlFieldsValidator()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errs := valuesValidator.Run(tc.krtYaml)
			if tc.wantError {
				require.True(t, len(errs) >= 1)
				assert.Error(t, errs[0])
				assert.Equal(t, tc.errorString, errs[0].Error())
				return
			}

			assert.Empty(t, errs)
		})
	}
}
