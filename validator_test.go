// //go:build unit

package main

// import (
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/konstellation-io/kai/engine/admin-api/mocks"
// 	"github.com/stretchr/testify/assert"
// )

// func TestKrtValidator_Run(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		krtYaml     *Krt
// 		wantError   bool
// 		errorString string
// 	}{
// 		{
// 			name:        "validation for a valid workflow",
// 			krtYaml:     NewKrtBuilder().Build(),
// 			wantError:   false,
// 			errorString: "",
// 		},
// 		{
// 			name: "fails if process name is not unique in workflow",
// 			krtYaml: NewKrtBuilder().
// 				WithProcesses([]Process{
// 					{
// 						Name:  "test-trigger",
// 						Type:  ProcessTypeTrigger,
// 						Build: ProcessBuild{Image: "test-trigger-image"},
// 						Subscriptions: []string{
// 							"test-exit",
// 						},
// 					},
// 					{
// 						Name:  "test-trigger",
// 						Type:  ProcessTypeTrigger,
// 						Build: ProcessBuild{Image: "test-trigger-image"},
// 						Subscriptions: []string{
// 							"test-exit",
// 						},
// 					},
// 				}).Build(),
// 			wantError:   true,
// 			errorString: ErrRepeatedProcessName.Error(),
// 		},
// 	}

// 	ctrl := gomock.NewController(t)
// 	fieldsValidator := mocks.NewMockFieldsValidator(ctrl)
// 	validator := NewKrtValidator(fieldsValidator)

// 	for _, tc := range tests {
// 		fieldsValidator.EXPECT().Run(tc.krtYaml).Return(nil)

// 		t.Run(tc.name, func(t *testing.T) {
// 			err := validator.Run(tc.krtYaml)
// 			if tc.wantError {
// 				assert.EqualError(t, err, tc.errorString)
// 				return
// 			}

// 			assert.NoError(t, err)
// 		})
// 		ctrl.Finish()
// 	}
// }
