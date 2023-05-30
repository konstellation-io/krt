//go:build unit

package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/konstellation-io/kai/engine/admin-api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestKrtValidator_Run(t *testing.T) {
	tests := []struct {
		name        string
		krtYaml     *Krt
		wantError   bool
		errorString string
	}{
		{
			name:        "validation for a valid workflow",
			krtYaml:     NewKrtBuilder().Build(),
			wantError:   false,
			errorString: "",
		},
		{
			name:        "fails if there are no wokflows",
			krtYaml:     NewKrtBuilder().WithWorkflows([]Workflow{}).Build(),
			wantError:   true,
			errorString: "process \"test-trigger\" requires at least one subscription",
		},
		{
			name: "fails if a process does not have any subscriptions",
			krtYaml: NewKrtBuilder().WithProcessesForWorkflow([]Process{
				{
					Name:  "test-trigger",
					Type:  ProcessTypeTrigger,
					Build: ProcessBuild{Image: "test-trigger-image"},
				},
			}, 0).Build(),
			wantError:   true,
			errorString: "process \"test-trigger\" requires at least one subscription",
		},
		{
			name: "fails if node name is not unique in workflow",
			krtYaml: NewKrtBuilder().
				WithProcessesForWorkflow([]Process{
					{
						Name:  "test-trigger",
						Type:  ProcessTypeTrigger,
						Build: ProcessBuild{Image: "test-trigger-image"},
						Subscriptions: []string{
							"test-exit",
						},
					},
					{
						Name:  "test-trigger",
						Type:  ProcessTypeTrigger,
						Build: ProcessBuild{Image: "test-trigger-image"},
						Subscriptions: []string{
							"test-exit",
						},
					},
				}, 0).Build(),
			wantError:   true,
			errorString: ErrRepeatedNodeName.Error(),
		},
	}

	ctrl := gomock.NewController(t)
	fieldsValidator := mocks.NewMockFieldsValidator(ctrl)
	validator := NewKrtValidator(fieldsValidator)

	for _, tc := range tests {
		fieldsValidator.EXPECT().Run(tc.krtYaml).Return(nil)

		t.Run(tc.name, func(t *testing.T) {
			err := validator.Run(tc.krtYaml)
			if tc.wantError {
				assert.EqualError(t, err, tc.errorString)
				return
			}

			assert.Empty(t, err)
		})
		ctrl.Finish()
	}
}
