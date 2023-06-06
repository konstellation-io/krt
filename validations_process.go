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

	if process.Subscriptions == nil || len(process.Subscriptions) == 0 {
		totalError = errors.MergeErrors(
			totalError,
			errors.MissingRequiredFieldError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions",
					workflowIdx,
					processIdx,
				),
			),
		)
	}

	err = process.validateNetworking(workflowIdx, processIdx)
	totalError = errors.MergeErrors(totalError, err)

	return totalError
}

func (process *Process) validateName(workflowIdx, processIdx int) error {
	return validateName(
		process.Name,
		fmt.Sprintf("krt.workflows[%d].processes[%d].name", workflowIdx, processIdx),
	)
}

func (process *Process) validateType(workflowIdx, processIdx int) error {
	if !isValidProcessType(string(process.Type)) {
		return errors.InvalidProcessTypeError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].type", workflowIdx, processIdx),
		)
	}

	return nil
}

func (process *Process) validateBuild(workflowIdx, processIdx int) error {
	if process.Build.Dockerfile == "" && process.Build.Image == "" {
		return errors.InvalidProcessBuildError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].build", workflowIdx, processIdx),
		)
	}

	return nil
}

func (process *Process) validateObjectStore(workflowIdx, processIdx int) error {
	EmptyObjectStore := ProcessObjectStore{}
	if process.ObjectStore == EmptyObjectStore {
		return nil
	}

	var totalError error
	if process.ObjectStore.Name == "" {
		totalError = errors.MergeErrors(
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
		totalError = errors.MergeErrors(totalError, err)
	}

	if !isValidObjectStoreScope(string(process.ObjectStore.Scope)) {
		totalError = errors.MergeErrors(
			totalError,
			errors.InvalidProcessObjectStoreScopeError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.scope", workflowIdx, processIdx),
			),
		)
	}

	return totalError
}

func (process *Process) validateNetworking(workflowIdx, processIdx int) error {
	emptyNetworking := ProcessNetworking{}
	if process.Networking == emptyNetworking {
		return nil
	}

	var totalError error
	if process.Networking.TargetPort == 0 {
		totalError = errors.MergeErrors(
			totalError,
			errors.MissingRequiredFieldError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].networking.targetPort", workflowIdx, processIdx),
			),
		)
	}

	if !isValidNetworkingProtocol(string(process.Networking.TargetProtocol)) {
		totalError = errors.MergeErrors(
			totalError, errors.InvalidNetworkingProtocolError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].networking.targetProtocol", workflowIdx, processIdx),
			),
		)
	}

	if process.Networking.DestinationPort == 0 {
		totalError = errors.MergeErrors(
			totalError,
			errors.MissingRequiredFieldError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].networking.destinationPort", workflowIdx, processIdx),
			),
		)
	}

	if !isValidNetworkingProtocol(string(process.Networking.DestinationProtocol)) {
		totalError = errors.MergeErrors(
			totalError,
			errors.InvalidNetworkingProtocolError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].networking.destinationProtocol", workflowIdx, processIdx),
			),
		)
	}

	return totalError
}

// validateSubscritpions checks if subscriptions for all process are valid.
// All requirements for subscritpions to be valid can be found in the readme.
func validateSubscritpions(subscriptions []Process, workflowIdx int) error {
	var (
		totalError          error
		processTypesByNames = make(map[string]ProcessType)
	)

	processCountByType := map[ProcessType]int{
		ProcessTypeTrigger: 0,
		ProcessTypeTask:    0,
		ProcessTypeExit:    0,
	}

	// loop 1, load processes names and type
	// also, checks if there are enough processes, a duplicated process name or duplicated subscriptions
	for processIdx, process := range subscriptions {
		var subscriptionAlreadyExists = make(map[string]bool)
		for _, subscription := range process.Subscriptions {
			if _, ok := subscriptionAlreadyExists[subscription]; ok {
				totalError = errors.MergeErrors(
					totalError,
					errors.DuplicatedProcessSubscriptionError(
						fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions.%s",
							workflowIdx,
							processIdx,
							subscription,
						),
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
			totalError = errors.MergeErrors(
				totalError,
				errors.DuplicatedProcessNameError(
					fmt.Sprintf("krt.workflows[%d].processes[%d].name", workflowIdx, processIdx),
				),
			)
		} else {
			processTypesByNames[process.Name] = process.Type
		}
	}

	if processCountByType[ProcessTypeTrigger] < 1 || processCountByType[ProcessTypeExit] < 1 {
		totalError = errors.MergeErrors(
			totalError,
			errors.NotEnoughProcessesError(
				fmt.Sprintf("krt.workflows[%d].processes", workflowIdx),
			),
		)
	}

	// loop 2, check if all subscriptions are valid
	for processIdx, process := range subscriptions {
		for _, subscription := range process.Subscriptions {
			if process.Name == subscription {
				totalError = errors.MergeErrors(totalError, errors.CannotSubscribeToItselfError(
					fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions.%s",
						workflowIdx,
						processIdx,
						subscription,
					),
				))
			}

			if !isValidSubscription(process.Type, processTypesByNames[subscription]) {
				totalError = errors.MergeErrors(totalError, errors.InvalidProcessSubscriptionError(
					string(process.Type),
					string(processTypesByNames[subscription]),
					fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions.%s",
						workflowIdx,
						processIdx,
						subscription,
					),
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