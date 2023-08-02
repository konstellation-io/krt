package krt

import (
	"fmt"

	"github.com/konstellation-io/krt/pkg/errors"
)

func (workflow *Workflow) Validate(workflowIdx int) error {
	return errors.Join(
		workflow.ValidateName(workflowIdx),
		workflow.ValidateType(workflowIdx),
		workflow.ValidateVersionConfig(workflowIdx),
		workflow.ValidateProcesses(workflowIdx),
	)
}

func (workflow *Workflow) ValidateName(workflowIdx int) error {
	return validateName(workflow.Name, fmt.Sprintf("krt.workflows[%d].name", workflowIdx))
}

func (workflow *Workflow) ValidateType(workflowIdx int) error {
	if !workflow.Type.IsValid() {
		return errors.InvalidWorkflowTypeError(fmt.Sprintf("krt.workflows[%d].type", workflowIdx))
	}

	return nil
}

func (workflow *Workflow) ValidateVersionConfig(workflowIdx int) error {
	return nil
}

func (workflow *Workflow) ValidateProcesses(workflowIdx int) error {
	var totalError error

	if len(workflow.Processes) == 0 {
		totalError = errors.Join(
			totalError,
			errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes", workflowIdx)),
		)
	} else {
		for idx, process := range workflow.Processes {
			totalError = errors.Join(totalError, process.Validate(workflowIdx, idx))
		}

		totalError = errors.Join(totalError, validateSubscritpions(workflow.Processes, workflowIdx))
	}

	return totalError
}

func validateWorkflowDuplicates(workflows []Workflow) error {
	var totalError error

	workflowNames := make(map[string]bool)
	for workflowIdx, workflow := range workflows {
		if workflowNames[workflow.Name] {
			totalError = errors.Join(
				totalError,
				errors.DuplicatedWorkflowNameError(
					fmt.Sprintf("krt.workflows[%d].name", workflowIdx),
				),
			)
		}

		workflowNames[workflow.Name] = true
	}

	return totalError
}
