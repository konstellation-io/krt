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
		totalError = errors.MergeErrors(totalError, errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions", workflowIdx, processIdx)))
	}

	err = process.validateNetworking(workflowIdx, processIdx)
	totalError = errors.MergeErrors(totalError, err)

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

func (process *Process) validateBuild(workflowIdx, processIdx int) error {
	if process.Build.Dockerfile == "" && process.Build.Image == "" {
		return errors.InvalidProcessBuildError(fmt.Sprintf("krt.workflows[%d].processes[%d].build", workflowIdx, processIdx))
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
		totalError = errors.MergeErrors(totalError, errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.name", workflowIdx, processIdx)))
	} else {
		err := validateName(process.ObjectStore.Name, fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.name", workflowIdx, processIdx))
		totalError = errors.MergeErrors(totalError, err)
	}

	if _, ok := ObjectStoreScopeMap[string(process.ObjectStore.Scope)]; !ok {
		totalError = errors.MergeErrors(totalError, errors.InvalidProcessObjectStoreScopeError(fmt.Sprintf("krt.workflows[%d].processes[%d].objectStore.scope", workflowIdx, processIdx)))
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
		totalError = errors.MergeErrors(totalError, errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes[%d].networking.targetPort", workflowIdx, processIdx)))
	}

	if _, ok := NetworkingProtocolMap[string(process.Networking.TargetProtocol)]; !ok {
		totalError = errors.MergeErrors(totalError, errors.InvalidNetworkingProtocolError(fmt.Sprintf("krt.workflows[%d].processes[%d].networking.targetProtocol", workflowIdx, processIdx)))
	}

	if process.Networking.DestinationPort == 0 {
		totalError = errors.MergeErrors(totalError, errors.MissingRequiredFieldError(fmt.Sprintf("krt.workflows[%d].processes[%d].networking.destinationPort", workflowIdx, processIdx)))
	}

	if _, ok := NetworkingProtocolMap[string(process.Networking.DestinationProtocol)]; !ok {
		totalError = errors.MergeErrors(totalError, errors.InvalidNetworkingProtocolError(fmt.Sprintf("krt.workflows[%d].processes[%d].networking.destinationProtocol", workflowIdx, processIdx)))
	}

	return totalError
}

// validateSubscritpions checks if subscriptions for all process are valid.
// All requirements for subscritpions to be valid can be found in the readme.
func validateSubscritpions(subscriptions []Process, workflowIdx int) error {
	var totalError error
	var processTypesByNames = make(map[string]ProcessType)

	// loop 1, load processes names and type
	// also, check if there are duplicated names
	for processIdx, process := range subscriptions {

		for _, subscription := range process.Subscriptions {
			var subscriptionAlreadyExists = make(map[string]bool)
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

		processTypesByNames[process.Name] = process.Type
	}

	// loop 2, check if all subscriptions are valid
	for processIdx, process := range subscriptions {
		for _, subscription := range process.Subscriptions {
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

// TODO:
// Subscriptions validation logic
// Github actions
// Sonarcloud config
// Parse methods
// Tests with yaml files
