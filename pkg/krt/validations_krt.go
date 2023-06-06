package krt

import "github.com/konstellation-io/krt/pkg/errors"

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
		err = validateWorkflowDuplicates(krt.Workflows)
		totalError = errors.MergeErrors(totalError, err)

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

// Things left:
// Parse methods
// Tests with yaml files
