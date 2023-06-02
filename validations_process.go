package main

import (
	"fmt"

	"github.com/konstellation-io/krt/errors"
)

func (process *Process) Validate(workflowIdx, processIdx int) error {
	var totalError error

	err := process.validateName(workflowIdx, processIdx)
	totalError = errors.MergeErrors(totalError, err)

	err = process.validateType(workflowIdx, processIdx)
	totalError = errors.MergeErrors(totalError, err)

	// err = process.validateBuild()
	// 	totalError = errors.MergeErrors(totalError, err)/ }

	// err = process.validateReplicas()
	// 	totalError = errors.MergeErrors(totalError, err)/ }

	// err = process.validateGPU()
	// 	totalError = errors.MergeErrors(totalError, err)/ }

	// err = process.validateObjectStore()
	// 	totalError = errors.MergeErrors(totalError, err)/ }

	// err = process.validateSecrets()
	// 	totalError = errors.MergeErrors(totalError, err)/ }

	if process.Subscriptions == nil || len(process.Subscriptions) == 0 {
		totalError = errors.MergeErrors(totalError, errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions", workflowIdx, processIdx)))
	}

	// err = process.validateSubscriptions()
	// 	totalError = errors.MergeErrors(totalError, err)/ }

	// err = process.validateNetworking()
	// 	totalError = errors.MergeErrors(totalError, err)/ }

	return totalError
}

func (process *Process) validateName(workflowIdx, processIdx int) error {
	return validateName(process.Name, fmt.Sprintf("krt.workflows[%d].processes[%d].name", workflowIdx, processIdx))
}

func (process *Process) validateType(workflowIdx, processIdx int) error {
	if _, ok := ProcessTypeMap[string(process.Type)]; !ok {
		return errors.InvalidProcessTypeError(fmt.Sprintf("krt.workflows[%d].processes[%d].type", workflowIdx, processIdx))
	}
	return nil
}
