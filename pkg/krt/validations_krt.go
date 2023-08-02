package krt

import "github.com/konstellation-io/krt/pkg/errors"

func (krt *Krt) Validate() error {
	return errors.Join(
		krt.ValidateDescription(),
		krt.ValidateKRTVersion(),
		krt.ValidateVersionConfig(),
		krt.ValidateWorkflows(),
	)
}

func (krt *Krt) ValidateDescription() error {
	if krt.Description == "" {
		return errors.MissingRequiredFieldError("krt.description")
	}

	return nil
}

func (krt *Krt) ValidateKRTVersion() error {
	return validateVersion(krt.Version, "krt.version")
}

func (krt *Krt) ValidateVersionConfig() error {
	return nil
}

func (krt *Krt) ValidateWorkflows() error {
	var totalError error

	if len(krt.Workflows) == 0 {
		totalError = errors.Join(
			totalError,
			errors.MissingRequiredFieldError("krt.workflows"),
		)
	} else {
		err := validateWorkflowDuplicates(krt.Workflows)
		totalError = errors.Join(totalError, err)

		for idx, workflow := range krt.Workflows {
			err := workflow.Validate(idx)
			totalError = errors.Join(totalError, err)
		}
	}

	return totalError
}
