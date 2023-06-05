package main

import (
	"fmt"

	"github.com/konstellation-io/krt/errors"
)

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

		err := validateSubscritpions(workflow.Processes, workflowIdx)
		totalError = errors.MergeErrors(totalError, err)
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

func validateWorkflowDuplicates(workflows []Workflow) error {
	var totalError error

	workflowNames := make(map[string]bool)
	for workflowIdx, workflow := range workflows {
		if workflowNames[workflow.Name] {
			totalError = errors.MergeErrors(totalError, errors.DuplicatedWorkflowNameError(fmt.Sprintf("krt.workflows[%d].name", workflowIdx)))
		}
		workflowNames[workflow.Name] = true
	}

	return totalError
}
