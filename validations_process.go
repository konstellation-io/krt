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

	err = process.validateBuild(workflowIdx, processIdx)
	totalError = errors.MergeErrors(totalError, err)

	err = process.validateObjectStore(workflowIdx, processIdx)
	totalError = errors.MergeErrors(totalError, err)

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

func (process *Process) validateBuild(workflowIdx, processIdx int) error {
	if process.Build.Dockerfile == "" && process.Build.Image == "" {
		return errors.InvalidProcessBuildError(fmt.Sprintf("krt.workflows[%d].processes[%d].build", workflowIdx, processIdx))
	}
	return nil
}

func (process *Process) validateObjectStore(workflowIdx, processIdx int) error {
	EmptyObjectStore := ProcessObjectStore{}
	if process.ObjectStore == EmptyObjectStore {
		return nil
	}

	if process.ObjectStore.Name == "" {
		return errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.name", workflowIdx, processIdx))
	}

	return validateName(process.ObjectStore.Name, fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.name", workflowIdx, processIdx))
}
