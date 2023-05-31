package main

import (
	"fmt"
)

//go:generate mockgen -source=${GOFILE} -destination=../../../../mocks/validator_${GOFILE} -package=mocks

type Validator interface {
	Run(krt *Krt) error
}

type KrtValidator struct {
	fieldsValidator FieldsValidator
}

func NewKrtValidator(fieldsValidator FieldsValidator) Validator {
	return &KrtValidator{
		fieldsValidator: fieldsValidator,
	}
}

func (v *KrtValidator) Run(krtYaml *Krt) error {
	var errs []error

	fieldValidationErrors := v.fieldsValidator.Run(krtYaml)

	if fieldValidationErrors != nil {
		errs = append(errs, fieldValidationErrors...)
	}

	workflowValidationErrors := v.getWorkflowsValidationErrors(krtYaml.Workflows)

	if workflowValidationErrors != nil {
		errs = append(errs, workflowValidationErrors...)
	}

	if errs != nil {
		return NewValidationError(errs)
	}

	return nil
}

func (v *KrtValidator) getWorkflowsValidationErrors(workflows []Workflow) []error {
	var validationErrors []error

	for _, workflow := range workflows {
		existingProcesses := make(map[string]bool, len(workflow.Processes))
		for _, process := range workflow.Processes {
			processNameAlreadyInUse := existingProcesses[process.Name]

			if processNameAlreadyInUse {
				validationErrors = append(validationErrors, ErrRepeatedProcessName)
			}

			existingProcesses[process.Name] = true

			if len(process.Subscriptions) < 1 {
				//nolint:goerr113 // errors need to be dynamically generated
				validationErrors = append(validationErrors, fmt.Errorf("process %q requires at least one subscription", process.Name))
			}
		}
	}

	return validationErrors
}
