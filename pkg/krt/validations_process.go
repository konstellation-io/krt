package krt

import (
	"fmt"

	"github.com/konstellation-io/krt/pkg/errors"
)

const subscritpionLocation = "krt.workflows[%d].processes[%d].subscriptions.%s"

func (process *Process) Validate(workflowIdx, processIdx int) error {
	return errors.Join(
		process.validateName(workflowIdx, processIdx),
		process.validateType(workflowIdx, processIdx),
		process.validateImage(workflowIdx, processIdx),
		process.validateReplicas(workflowIdx, processIdx),
		process.validateGPU(workflowIdx, processIdx),
		process.validateConfig(workflowIdx, processIdx),
		process.validateObjectStore(workflowIdx, processIdx),
		process.validateSecrets(workflowIdx, processIdx),
		process.validateSubscriptions(workflowIdx, processIdx),
		process.validateNetworking(workflowIdx, processIdx),
	)
}

func (process *Process) validateName(workflowIdx, processIdx int) error {
	return validateName(
		process.Name,
		fmt.Sprintf("krt.workflows[%d].processes[%d].name", workflowIdx, processIdx),
	)
}

func (process *Process) validateType(workflowIdx, processIdx int) error {
	if !process.Type.IsValid() {
		return errors.InvalidProcessTypeError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].type", workflowIdx, processIdx),
		)
	}

	return nil
}

func (process *Process) validateImage(workflowIdx, processIdx int) error {
	if process.Image == "" {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].image", workflowIdx, processIdx),
		)
	}

	return nil
}

func (process *Process) validateReplicas(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) validateGPU(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) validateConfig(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) validateObjectStore(workflowIdx, processIdx int) error {
	if process.ObjectStore == nil {
		return nil
	}

	var totalError error
	if process.ObjectStore.Name == "" {
		totalError = errors.Join(
			totalError,
			errors.MissingRequiredFieldError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.name", workflowIdx, processIdx),
			),
		)
	} else {
		err := validateName(
			process.ObjectStore.Name,
			fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.name", workflowIdx, processIdx),
		)
		totalError = errors.Join(totalError, err)
	}

	if !process.ObjectStore.Scope.IsValid() {
		totalError = errors.Join(
			totalError,
			errors.InvalidProcessObjectStoreScopeError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.scope", workflowIdx, processIdx),
			),
		)
	}

	return totalError
}

func (process *Process) validateSecrets(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) validateSubscriptions(workflowIdx, processIdx int) error {
	if process.Subscriptions == nil {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions",
				workflowIdx,
				processIdx,
			),
		)
	}

	return nil
}

func (process *Process) validateNetworking(workflowIdx, processIdx int) error {
	if process.Networking == nil {
		return nil
	}

	var totalError error
	if process.Networking.TargetPort == 0 {
		totalError = errors.Join(
			totalError,
			errors.MissingRequiredFieldError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].networking.targetPort", workflowIdx, processIdx),
			),
		)
	}

	if process.Networking.DestinationPort == 0 {
		totalError = errors.Join(
			totalError,
			errors.MissingRequiredFieldError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].networking.destinationPort", workflowIdx, processIdx),
			),
		)
	}

	if !process.Networking.Protocol.IsValid() {
		totalError = errors.Join(
			totalError, errors.InvalidNetworkingProtocolError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].networking.protocol", workflowIdx, processIdx),
			),
		)
	}

	return totalError
}

// validateSubscritpions checks if subscriptions for all process are valid.
// All requirements for subscritpions to be valid can be found in the readme.
func validateSubscritpions(processes []Process, workflowIdx int) error {
	var totalError error

	processTypesByNames, err := countProcessesSubscriptions(processes, workflowIdx)
	totalError = errors.Join(totalError, err)

	totalError = errors.Join(totalError, checkSubscriptions(processes, workflowIdx, processTypesByNames))

	return totalError
}

// countProcessesSubscriptions, will load processes types by their names
// also, checks if there are enough processes, a duplicated process name or duplicated subscriptions.
func countProcessesSubscriptions(processes []Process, workflowIdx int) (map[string]ProcessType, error) {
	var (
		totalError          error
		processTypesByNames = make(map[string]ProcessType)
	)

	processCountByType := map[ProcessType]int{
		ProcessTypeTrigger: 0,
		ProcessTypeTask:    0,
		ProcessTypeExit:    0,
	}

	for processIdx, process := range processes {
		var subscriptionAlreadyExists = make(map[string]bool)
		for _, subscription := range process.Subscriptions {
			if _, ok := subscriptionAlreadyExists[subscription]; ok {
				totalError = errors.Join(
					totalError,
					errors.DuplicatedProcessSubscriptionError(
						fmt.Sprintf(subscritpionLocation, workflowIdx, processIdx, subscription),
					),
				)
			} else {
				subscriptionAlreadyExists[subscription] = true
			}
		}

		if _, ok := processCountByType[process.Type]; ok {
			processCountByType[process.Type]++
		}

		if _, ok := processTypesByNames[process.Name]; ok {
			totalError = errors.Join(
				totalError,
				errors.DuplicatedProcessNameError(
					fmt.Sprintf("krt.workflows[%d].processes[%d].name", workflowIdx, processIdx),
				),
			)
		} else {
			processTypesByNames[process.Name] = process.Type
		}
	}

	totalError = errors.Join(totalError, checkProcessCount(processCountByType, workflowIdx))

	return processTypesByNames, totalError
}

func checkProcessCount(processesCount map[ProcessType]int, workflowIdx int) error {
	if processesCount[ProcessTypeTrigger] < 1 || processesCount[ProcessTypeExit] < 1 {
		return errors.NotEnoughProcessesError(
			fmt.Sprintf("krt.workflows[%d].processes", workflowIdx),
		)
	}

	return nil
}

func checkSubscriptions(processes []Process, workflowIdx int, processTypesByNames map[string]ProcessType) error {
	var totalError error

	for processIdx, process := range processes {
		for _, subscription := range process.Subscriptions {
			if process.Name == subscription {
				totalError = errors.Join(totalError, errors.CannotSubscribeToItselfError(
					fmt.Sprintf(subscritpionLocation, workflowIdx, processIdx, subscription),
				))

				continue
			}

			fetchedSubscription, processExists := processTypesByNames[subscription]
			if !processExists {
				totalError = errors.Join(totalError, errors.CannotSubscribeToNonExistentProcessError(
					subscription,
					fmt.Sprintf("krt.workflows[%d].processes[%d]", workflowIdx, processIdx),
				))

				continue
			}

			if !isValidSubscription(process.Type, fetchedSubscription) {
				totalError = errors.Join(totalError, errors.InvalidProcessSubscriptionError(
					string(process.Type),
					string(fetchedSubscription),
					fmt.Sprintf(subscritpionLocation, workflowIdx, processIdx, subscription),
				))
			}
		}
	}

	return totalError
}

func isValidSubscription(processType, subscriptionProcessType ProcessType) bool {
	switch processType {
	case ProcessTypeTrigger:
		return subscriptionProcessType == ProcessTypeExit
	case ProcessTypeTask, ProcessTypeExit:
		return subscriptionProcessType != ProcessTypeExit
	default:
		return false
	}
}
