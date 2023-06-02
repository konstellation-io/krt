package main

import (
	"fmt"
	"regexp"

	"github.com/konstellation-io/krt/errors"
)

const maxFieldNameLength = 20

func (krt *Krt) Validate() []error {
	var validationErrors []error

	err := krt.validateName()
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	err = krt.validateDescription()
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	err = krt.validateVersion()
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	if len(krt.Workflows) == 0 {
		validationErrors = append(
			validationErrors,
			errors.MissingRequiredFieldError("krt.workflows"),
		)
	} else {
		for idx, workflow := range krt.Workflows {
			errs := workflow.Validate(idx)
			if errs != nil {
				validationErrors = append(validationErrors, errs...)
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
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

func (workflow *Workflow) Validate(workflowIdx int) []error {
	var validationErrors []error

	err := workflow.validateName(workflowIdx)
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	err = workflow.validateType(workflowIdx)
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	if len(workflow.Processes) == 0 {
		validationErrors = append(
			validationErrors,
			errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes", workflowIdx)),
		)
	} else {
		for idx, process := range workflow.Processes {
			errs := process.Validate(workflowIdx, idx)
			if errs != nil {
				validationErrors = append(validationErrors, errs...)
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
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

func (process *Process) Validate(workflowIdx, processIdx int) []error {
	var validationErrors []error

	err := process.validateName(workflowIdx, processIdx)
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	err = process.validateType(workflowIdx, processIdx)
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	// err = process.validateBuild()
	// if err != nil {
	// 	validationErrors = append(validationErrors, err)
	// }

	// err = process.validateReplicas()
	// if err != nil {
	// 	validationErrors = append(validationErrors, err)
	// }

	// err = process.validateGPU()
	// if err != nil {
	// 	validationErrors = append(validationErrors, err)
	// }

	// err = process.validateObjectStore()
	// if err != nil {
	// 	validationErrors = append(validationErrors, err)
	// }

	// err = process.validateSecrets()
	// if err != nil {
	// 	validationErrors = append(validationErrors, err)
	// }

	if process.Subscriptions == nil || len(process.Subscriptions) == 0 {
		validationErrors = append(validationErrors, errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions", workflowIdx, processIdx)))
	}

	// err = process.validateSubscriptions()
	// if err != nil {
	// 	validationErrors = append(validationErrors, err)
	// }

	// err = process.validateNetworking()
	// if err != nil {
	// 	validationErrors = append(validationErrors, err)
	// }

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
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
