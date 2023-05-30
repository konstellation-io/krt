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
		for _, node := range workflow.Processes {
			nodeNameAlreadyInUse := existingProcesses[node.Name]

			if nodeNameAlreadyInUse {
				validationErrors = append(validationErrors, ErrRepeatedNodeName)
			}

			existingProcesses[node.Name] = true

			if len(node.Subscriptions) < 1 {
				//nolint:goerr113 // errors need to be dynamically generated
				validationErrors = append(validationErrors, fmt.Errorf("node %q requires at least one subscription", node.Name))
			}
		}
	}

	return validationErrors
}
