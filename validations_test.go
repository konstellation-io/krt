//go:build unit

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konstellation-io/krt/errors"
)

const largeName = "this-name-is-higher-than-the-maximum-allowed"
const invalidName = "Invalid string!"

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
			name:        "fails if krt hasn't required process build",
			krtYaml:     NewKrtBuilder().WithProcessBuild(ProcessBuild{}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessBuild,
			errorString: errors.InvalidProcessBuildError("krt.workflows[0].processes[0].build").Error(),
		},
		{
			name: "fails if krt hasn't required object store name if declared",
			krtYaml: NewKrtBuilder().WithProcessObjectStore(
				ProcessObjectStore{
					Name:  "",
					Scope: ObjectStoreScopeProduct,
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].objectStore.name").Error(),
		},
		{
			name:        "fails if krt hasn't required process subscriptions",
			krtYaml:     NewKrtBuilder().WithProcessSubscriptions(nil, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].subscriptions").Error(),
		},
		{
			name:        "fails if krt required process subscriptions are empty",
			krtYaml:     NewKrtBuilder().WithProcessSubscriptions([]string{}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].subscriptions").Error(),
		},
		{
			name: "fails if krt hasn't required networking target port if declared",
			krtYaml: NewKrtBuilder().WithProcessNetworking(
				ProcessNetworking{
					TargetProtocol:      "UDP",
					DestinationPort:     9000,
					DestinationProtocol: "UDP",
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].networking.targetPort").Error(),
		},
		{
			name: "fails if krt hasn't required networking destination port if declared",
			krtYaml: NewKrtBuilder().WithProcessNetworking(
				ProcessNetworking{
					TargetPort:          9000,
					TargetProtocol:      "UDP",
					DestinationProtocol: "UDP",
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].networking.destinationPort").Error(),
		},
	}

	invalidNameTests := []test{
		{
			name:        "fails if version name has an invalid format",
			krtYaml:     NewKrtBuilder().WithVersion(invalidName).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidFieldName,
			errorString: errors.InvalidFieldNameError("krt.version").Error(),
		},
		{
			name:        "fails if version name has an invalid length",
			krtYaml:     NewKrtBuilder().WithVersion(largeName).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.version", maxFieldNameLength).Error(),
		},
		{
			name:        "fails if krt workflow name has an invalid format",
			krtYaml:     NewKrtBuilder().WithWorkflowName(invalidName).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidFieldName,
			errorString: errors.InvalidFieldNameError("krt.workflows[0].name").Error(),
		},
		{
			name:        "fails if krt workflow name has an invalid length",
			krtYaml:     NewKrtBuilder().WithWorkflowName(largeName).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].name", maxFieldNameLength).Error(),
		},
		{
			name: "fails if krt workflow name is duplicated",
			krtYaml: NewKrtBuilder().WithWorkflows([]Workflow{
				{
					Name:      "test-workflow",
					Type:      WorkflowTypeTraining,
					Processes: []Process{},
				},
				{
					Name:      "test-workflow",
					Type:      WorkflowTypeTraining,
					Processes: []Process{},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrDuplicatedWorkflowName,
			errorString: errors.DuplicatedWorkflowNameError("krt.workflows[1].name").Error(),
		},
		{
			name:        "fails if krt process name has an invalid format",
			krtYaml:     NewKrtBuilder().WithProcessName(invalidName, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidFieldName,
			errorString: errors.InvalidFieldNameError("krt.workflows[0].processes[0].name").Error(),
		},
		{
			name:        "fails if krt process name has an invalid length",
			krtYaml:     NewKrtBuilder().WithProcessName(largeName, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].processes[0].name", maxFieldNameLength).Error(),
		},
		{
			name: "fails if krt process object store name has an invalid format",
			krtYaml: NewKrtBuilder().WithProcessObjectStore(
				ProcessObjectStore{
					Name:  invalidName,
					Scope: ObjectStoreScopeProduct,
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidFieldName,
			errorString: errors.InvalidFieldNameError("krt.workflows[0].processes[0].objectStore.name").Error(),
		},
		{
			name: "fails if krt process object store name has an invalid length",
			krtYaml: NewKrtBuilder().WithProcessObjectStore(
				ProcessObjectStore{
					Name:  largeName,
					Scope: ObjectStoreScopeProduct,
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].processes[0].objectStore.name", maxFieldNameLength).Error(),
		},
		{
			name: "fails if krt process name is duplicated",
			krtYaml: NewKrtBuilder().WithProcesses([]Process{
				{
					Name:  "test-process",
					Type:  ProcessTypeTrigger,
					Build: ProcessBuild{Image: "test-image"},
				},
				{
					Name:  "test-process",
					Type:  ProcessTypeTask,
					Build: ProcessBuild{Image: "test-image"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrDuplicatedProcessName,
			errorString: errors.DuplicatedProcessNameError("krt.workflows[0].processes[1].name").Error(),
		},
	}

	invalidTypeTests := []test{
		{
			name:        "fails if krt hasn't a valid workflow type",
			krtYaml:     NewKrtBuilder().WithWorkflowType("invalid").Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidWorkflowType,
			errorString: errors.InvalidWorkflowTypeError("krt.workflows[0].type").Error(),
		},
		{
			name:        "fails if krt hasn't a valid process type",
			krtYaml:     NewKrtBuilder().WithProcessType("invalid", 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessType,
			errorString: errors.InvalidProcessTypeError("krt.workflows[0].processes[0].type").Error(),
		},
		{
			name: "fails if krt hasn't a valid process object store scope",
			krtYaml: NewKrtBuilder().WithProcessObjectStore(
				ProcessObjectStore{
					Name:  "test",
					Scope: "invalid",
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessObjectStoreScope,
			errorString: errors.InvalidProcessObjectStoreScopeError("krt.workflows[0].processes[0].objectStore.scope").Error(),
		},
		{
			name: "fails if krt hasn't a valid process networking target protocol",
			krtYaml: NewKrtBuilder().WithProcessNetworking(
				ProcessNetworking{
					TargetPort:          9000,
					TargetProtocol:      invalidName,
					DestinationPort:     9000,
					DestinationProtocol: "UDP",
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidNetworkingProtocol,
			errorString: errors.InvalidNetworkingProtocolError("krt.workflows[0].processes[0].networking.targetProtocol").Error(),
		},
		{
			name: "fails if krt hasn't a valid process networking destination protocol",
			krtYaml: NewKrtBuilder().WithProcessNetworking(
				ProcessNetworking{
					TargetPort:          9000,
					TargetProtocol:      "UDP",
					DestinationPort:     9000,
					DestinationProtocol: invalidName,
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidNetworkingProtocol,
			errorString: errors.InvalidNetworkingProtocolError("krt.workflows[0].processes[0].networking.destinationProtocol").Error(),
		},
	}

	invalidSubscriptionTests := []test{
		{
			name: "fails if krt has not enough processes",
			krtYaml: NewKrtBuilder().WithProcesses([]Process{
				{
					Name:          "test-trigger",
					Type:          ProcessTypeTrigger,
					Build:         ProcessBuild{Image: "test-trigger-image"},
					Subscriptions: []string{"test-task-1"},
				},
				{
					Name:          "test-task-1",
					Type:          ProcessTypeTask,
					Build:         ProcessBuild{Image: "test-task-image"},
					Subscriptions: []string{"test-task-2"},
				},
				{
					Name:          "test-task-2",
					Type:          ProcessTypeTask,
					Build:         ProcessBuild{Image: "test-task-image"},
					Subscriptions: []string{"test-task-1"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrNotEnoughProcesses,
			errorString: errors.NotEnoughProcessesError("krt.workflows[0].processes").Error(),
		},
		{
			name: "fails if krt has duplicated process subscriptions",
			krtYaml: NewKrtBuilder().WithProcesses([]Process{
				{
					Name:          "test-trigger",
					Type:          ProcessTypeTrigger,
					Build:         ProcessBuild{Image: "test-trigger-image"},
					Subscriptions: []string{"test-exit", "test-exit"},
				},
				{
					Name:          "test-exit",
					Type:          ProcessTypeExit,
					Build:         ProcessBuild{Image: "test-exit-image"},
					Subscriptions: []string{"test-trigger"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrDuplicatedProcessSubscription,
			errorString: errors.DuplicatedProcessSubscriptionError("krt.workflows[0].processes[0].subscriptions.test-exit").Error(),
		},
		{
			name: "fails if krt has invalid process subscriptions",
			krtYaml: NewKrtBuilder().WithProcesses([]Process{
				{
					Name:          "test-trigger",
					Type:          ProcessTypeTrigger,
					Build:         ProcessBuild{Image: "test-trigger-image"},
					Subscriptions: []string{"test-exit", "test-task"},
				},
				{
					Name:          "test-task",
					Type:          ProcessTypeTask,
					Build:         ProcessBuild{Image: "test-task-image"},
					Subscriptions: []string{"test-trigger"},
				},
				{
					Name:          "test-exit",
					Type:          ProcessTypeExit,
					Build:         ProcessBuild{Image: "test-exit-image"},
					Subscriptions: []string{"test-trigger"},
				},
			}).Build(),
			wantError: true,
			errorType: errors.ErrInvalidProcessSubscription,
			errorString: errors.InvalidProcessSubscriptionError(
				string(ProcessTypeTrigger),
				string(ProcessTypeTask),
				"krt.workflows[0].processes[0].subscriptions",
			).Error(),
		},
		{
			name: "fails if krt has a process subscribing to itself",
			krtYaml: NewKrtBuilder().WithProcesses([]Process{
				{
					Name:          "test-trigger",
					Type:          ProcessTypeTrigger,
					Build:         ProcessBuild{Image: "test-trigger-image"},
					Subscriptions: []string{"test-trigger"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrCannotSubscribeToItself,
			errorString: errors.CannotSubscribeToItselfError("krt.workflows[0].processes[0].subscriptions.test-trigger").Error(),
		},
	}

	allTests := make([]test, 0)
	allTests = append(allTests, correctBuildTests...)
	allTests = append(allTests, requiredFieldsTests...)
	allTests = append(allTests, invalidNameTests...)
	allTests = append(allTests, invalidTypeTests...)
	allTests = append(allTests, invalidSubscriptionTests...)

	for _, tc := range allTests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.krtYaml.Validate()
			if tc.wantError {
				assert.True(t, errors.Is(err, tc.errorType))
				assert.ErrorContains(t, err, tc.errorString)
			} else {
				assert.Empty(t, err)
			}
		})
	}
}
