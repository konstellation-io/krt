//go:build unit

package krt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
)

const largeName = "this-name-is-higher-than-the-maximum-allowed"
const invalidName = "Invalid string!"
const invalidVersion = "1.0.0"

type test struct {
	name        string
	krtYaml     *krt.Krt
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
			name:        "fails if krt hasn't required field version",
			krtYaml:     NewKrtBuilder().WithVersion("").Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.version").Error(),
		},
		{
			name:        "fails if krt hasn't required field description",
			krtYaml:     NewKrtBuilder().WithDescription("").Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.description").Error(),
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
			name:        "fails if krt hasn't required process image",
			krtYaml:     NewKrtBuilder().WithProcessImage("", 0).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].image").Error(),
		},
		{
			name: "fails if krt hasn't required object store name if declared",
			krtYaml: NewKrtBuilder().WithProcessObjectStore(
				&krt.ProcessObjectStore{
					Name:  "",
					Scope: krt.ObjectStoreScopeProduct,
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
			name: "fails if krt hasn't required networking target port if declared",
			krtYaml: NewKrtBuilder().WithProcessNetworking(
				&krt.ProcessNetworking{
					DestinationPort: 9000,
					Protocol:        "UDP",
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
				&krt.ProcessNetworking{
					TargetPort: 9000,
					Protocol:   "UDP",
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].networking.destinationPort").Error(),
		},
		{
			name: "fails if krt hasn't required process limit resources",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				nil,
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].resourceLimits").Error(),
		},
		{
			name: "fails if krt hasn't required process cpu",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					nil,
					&krt.ResourceLimit{
						Request: "100M",
						Limit:   "200M",
					},
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].resourceLimits.CPU").Error(),
		},
		{
			name: "fails if krt hasn't required process cpu request",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Limit: "1",
					},
					&krt.ResourceLimit{
						Request: "100M",
						Limit:   "200M",
					},
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].resourceLimits.CPU.request").Error(),
		},
		{
			name: "fails if krt hasn't required process memory",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "100m",
						Limit:   "200m",
					},
					nil,
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].resourceLimits.memory").Error(),
		},
		{
			name: "fails if krt hasn't required process memory request",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "100m",
						Limit:   "200m",
					},
					&krt.ResourceLimit{
						Limit: "100M",
					},
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrMissingRequiredField,
			errorString: errors.MissingRequiredFieldError("krt.workflows[0].processes[0].resourceLimits.memory.request").Error(),
		},
	}

	invalidNameTests := []test{
		{
			name:        "fails if version tag has an invalid format",
			krtYaml:     NewKrtBuilder().WithVersion(invalidVersion).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidVersionTag,
			errorString: errors.InvalidVersionTagError("krt.version").Error(),
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
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].name", krt.MaxFieldNameLength).Error(),
		},
		{
			name: "fails if krt workflow name is duplicated",
			krtYaml: NewKrtBuilder().WithWorkflows([]krt.Workflow{
				{
					Name:      "test-workflow",
					Type:      krt.WorkflowTypeTraining,
					Processes: []krt.Process{},
				},
				{
					Name:      "test-workflow",
					Type:      krt.WorkflowTypeTraining,
					Processes: []krt.Process{},
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
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].processes[0].name", krt.MaxFieldNameLength).Error(),
		},
		{
			name: "fails if krt process object store name has an invalid format",
			krtYaml: NewKrtBuilder().WithProcessObjectStore(
				&krt.ProcessObjectStore{
					Name:  invalidName,
					Scope: krt.ObjectStoreScopeProduct,
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
				&krt.ProcessObjectStore{
					Name:  largeName,
					Scope: krt.ObjectStoreScopeProduct,
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidLengthField,
			errorString: errors.InvalidLengthFieldError("krt.workflows[0].processes[0].objectStore.name", krt.MaxFieldNameLength).Error(),
		},
		{
			name: "fails if krt process name is duplicated",
			krtYaml: NewKrtBuilder().WithProcesses([]krt.Process{
				{
					Name:  "test-process",
					Type:  krt.ProcessTypeTrigger,
					Image: "test-image",
				},
				{
					Name:  "test-process",
					Type:  krt.ProcessTypeTask,
					Image: "test-image",
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrDuplicatedProcessName,
			errorString: errors.DuplicatedProcessNameError("krt.workflows[0].processes[1].name").Error(),
		},
		{
			name: "fails if krt cpu request has an invalid format",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "invalid",
						Limit:   "200m",
					},
					&krt.ResourceLimit{
						Request: "100M",
						Limit:   "200M",
					},
				}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessCPUResourceLimit,
			errorString: errors.InvalidProcessCPUError("krt.workflows[0].processes[0].resourceLimits.CPU.request").Error(),
		},
		{
			name: "fails if krt cpu limit has an invalid format",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "200m",
						Limit:   "invalid",
					},
					&krt.ResourceLimit{
						Request: "100M",
						Limit:   "200M",
					},
				}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessCPUResourceLimit,
			errorString: errors.InvalidProcessCPUError("krt.workflows[0].processes[0].resourceLimits.CPU.limit").Error(),
		},
		{
			name: "fails if krt memory request has an invalid format",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "200m",
						Limit:   "200m",
					},
					&krt.ResourceLimit{
						Request: "invalid",
						Limit:   "200m",
					},
				}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessMemoryResourceLimit,
			errorString: errors.InvalidProcessMemoryError("krt.workflows[0].processes[0].resourceLimits.memory.request").Error(),
		},
		{
			name: "fails if krt memory limit has an invalid format",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "200m",
						Limit:   "200m",
					},
					&krt.ResourceLimit{
						Request: "200m",
						Limit:   "invalid",
					},
				}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessMemoryResourceLimit,
			errorString: errors.InvalidProcessMemoryError("krt.workflows[0].processes[0].resourceLimits.memory.limit").Error(),
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
				&krt.ProcessObjectStore{
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
			name: "fails if krt hasn't a valid process networking protocol",
			krtYaml: NewKrtBuilder().WithProcessNetworking(
				&krt.ProcessNetworking{
					TargetPort:      9000,
					DestinationPort: 9000,
					Protocol:        invalidName,
				},
				0,
			).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidNetworkingProtocol,
			errorString: errors.InvalidNetworkingProtocolError("krt.workflows[0].processes[0].networking.protocol").Error(),
		},
	}

	invalidResourceRelationTests := []test{
		{
			name: "fails if krt cpu limit is lower than request",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "200m",
						Limit:   "0.1",
					},
					&krt.ResourceLimit{
						Request: "100M",
						Limit:   "200M",
					},
				}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessCPURelation,
			errorString: errors.InvalidProcessCPURelationError("krt.workflows[0].processes[0].resourceLimits.CPU").Error(),
		},
		{
			name: "fails if krt memory limit is lower than request",
			krtYaml: NewKrtBuilder().WithProcessResourceLimits(
				&krt.ProcessResourceLimits{
					&krt.ResourceLimit{
						Request: "200m",
						Limit:   "200m",
					},
					&krt.ResourceLimit{
						Request: "2Mi",
						Limit:   "2000k",
					}}, 0).Build(),
			wantError:   true,
			errorType:   errors.ErrInvalidProcessMemoryRelation,
			errorString: errors.InvalidProcessMemoryRelationError("krt.workflows[0].processes[0].resourceLimits.memory").Error(),
		},
	}

	invalidSubscriptionTests := []test{
		{
			name: "fails if krt has not enough processes",
			krtYaml: NewKrtBuilder().WithProcesses([]krt.Process{
				{
					Name:          "test-trigger",
					Type:          krt.ProcessTypeTrigger,
					Image:         "test-trigger-image",
					Subscriptions: []string{"test-task-1"},
				},
				{
					Name:          "test-task-1",
					Type:          krt.ProcessTypeTask,
					Image:         "test-task-image",
					Subscriptions: []string{"test-task-2"},
				},
				{
					Name:          "test-task-2",
					Type:          krt.ProcessTypeTask,
					Image:         "test-task-image",
					Subscriptions: []string{"test-task-1"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrNotEnoughProcesses,
			errorString: errors.NotEnoughProcessesError("krt.workflows[0].processes").Error(),
		},
		{
			name: "fails if krt has duplicated process subscriptions",
			krtYaml: NewKrtBuilder().WithProcesses([]krt.Process{
				{
					Name:          "test-trigger",
					Type:          krt.ProcessTypeTrigger,
					Image:         "test-trigger-image",
					Subscriptions: []string{"test-exit", "test-exit"},
				},
				{
					Name:          "test-exit",
					Type:          krt.ProcessTypeExit,
					Image:         "test-exit-image",
					Subscriptions: []string{"test-trigger"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrDuplicatedProcessSubscription,
			errorString: errors.DuplicatedProcessSubscriptionError("krt.workflows[0].processes[0].subscriptions.test-exit").Error(),
		},
		{
			name: "fails if krt has invalid process subscriptions",
			krtYaml: NewKrtBuilder().WithProcesses([]krt.Process{
				{
					Name:          "test-trigger",
					Type:          krt.ProcessTypeTrigger,
					Image:         "test-trigger-image",
					Subscriptions: []string{"test-exit", "test-task"},
				},
				{
					Name:          "test-task",
					Type:          krt.ProcessTypeTask,
					Image:         "test-task-image",
					Subscriptions: []string{"test-trigger"},
				},
				{
					Name:          "test-exit",
					Type:          krt.ProcessTypeExit,
					Image:         "test-exit-image",
					Subscriptions: []string{"test-trigger"},
				},
			}).Build(),
			wantError: true,
			errorType: errors.ErrInvalidProcessSubscription,
			errorString: errors.InvalidProcessSubscriptionError(
				string(krt.ProcessTypeTrigger),
				string(krt.ProcessTypeTask),
				"krt.workflows[0].processes[0].subscriptions",
			).Error(),
		},
		{
			name: "fails if krt has a process subscribing to itself",
			krtYaml: NewKrtBuilder().WithProcesses([]krt.Process{
				{
					Name:          "test-trigger",
					Type:          krt.ProcessTypeTrigger,
					Image:         "test-trigger-image",
					Subscriptions: []string{"test-trigger"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrCannotSubscribeToItself,
			errorString: errors.CannotSubscribeToItselfError("krt.workflows[0].processes[0].subscriptions.test-trigger").Error(),
		},
		{
			name: "fails if krt has a process subscribing to a non existing process",
			krtYaml: NewKrtBuilder().WithProcesses([]krt.Process{
				{
					Name:          "test-trigger",
					Type:          krt.ProcessTypeTrigger,
					Image:         "test-trigger-image",
					Subscriptions: []string{"non-existent"},
				},
			}).Build(),
			wantError:   true,
			errorType:   errors.ErrCannotSubscribeToNonExistentProcess,
			errorString: errors.CannotSubscribeToNonExistentProcessError("non-existent", "krt.workflows[0].processes[0]").Error(),
		},
	}

	allTests := make([]test, 0)
	allTests = append(allTests, correctBuildTests...)
	allTests = append(allTests, requiredFieldsTests...)
	allTests = append(allTests, invalidNameTests...)
	allTests = append(allTests, invalidTypeTests...)
	allTests = append(allTests, invalidResourceRelationTests...)
	allTests = append(allTests, invalidSubscriptionTests...)

	for _, tc := range allTests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.krtYaml.Validate()
			if tc.wantError {
				assert.ErrorIs(t, err, tc.errorType)
				assert.ErrorContains(t, err, tc.errorString)
			} else {
				assert.Empty(t, err)
			}
		})
	}
}
