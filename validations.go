package main

import (
	"fmt"
	"regexp"

	"github.com/konstellation-io/krt/errors"
)

const maxFieldNameLength = 20

func (krt *Krt) Validate() error {
	var totalError error

	err := krt.validateName()
	totalError = errors.MergeErrors(totalError, err)

	err = krt.validateDescription()
	totalError = errors.MergeErrors(totalError, err)

	err = krt.validateVersion()
	totalError = errors.MergeErrors(totalError, err)

	if len(krt.Workflows) == 0 {
		totalError = errors.MergeErrors(
			totalError,
			errors.MissingRequiredFieldError("krt.workflows"),
		)
	} else {
		for idx, workflow := range krt.Workflows {
			err := workflow.Validate(idx)
			totalError = errors.MergeErrors(totalError, err)
		}
	}

	return totalError
}

func (krt *Krt) validateName() error {
	return validateName(krt.Name, "krt.name")
}

func (krt *Krt) validateDescription() error {
	if krt.Description == "" {
		return errors.MissingRequiredFieldError("krt.description")
	}
	return nil
}

func (krt *Krt) validateVersion() error {
	return validateName(krt.Version, "krt.version")
}

func (workflow *Workflow) Validate(workflowIdx int) error {
	var totalError error

	err := workflow.validateName(workflowIdx)
	totalError = errors.MergeErrors(totalError, err)

	err = workflow.validateType(workflowIdx)
	totalError = errors.MergeErrors(totalError, err)

	if len(workflow.Processes) == 0 {
		totalError = errors.MergeErrors(
			totalError,
			errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes", workflowIdx)),
		)
	} else {
		for idx, process := range workflow.Processes {
			err := process.Validate(workflowIdx, idx)
			totalError = errors.MergeErrors(totalError, err)
		}
	}

	return totalError
}

func (workflow *Workflow) validateName(workflowIdx int) error {
	return validateName(workflow.Name, fmt.Sprintf("krt.workflows[%d].name", workflowIdx))
}

func (workflow *Workflow) validateType(workflowIdx int) error {
	if _, ok := WorkflowTypeMap[string(workflow.Type)]; !ok {
		return errors.InvalidWorkflowTypeError(fmt.Sprintf("krt.workflows[%d].type", workflowIdx))
	}
	return nil
}

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

func isValidResourceName(name string) bool {
	reResourceName := regexp.MustCompile("^[a-z0-9]([-a-z0-9]*[a-z0-9])?$")
	return reResourceName.MatchString(name)
}

func validateName(name, nameLocation string) error {
	if name == "" {
		return errors.MissingRequiredFieldError(nameLocation)
	}
	if !isValidResourceName(name) {
		return errors.InvalidFieldNameError(nameLocation)
	}
	if len(name) > maxFieldNameLength {
		return errors.InvalidLengthFieldError(nameLocation, maxFieldNameLength)
	}
	return nil
}
