package krt

import (
	"fmt"
	"strings"

	"github.com/konstellation-io/krt/pkg/errors"
)

const subscriptionLocation = "krt.workflows[%d].processes[%d].subscriptions.%s"

func (process *Process) Validate(workflowIdx, processIdx int) error {
	return errors.Join(
		process.ValidateName(workflowIdx, processIdx),
		process.ValidateType(workflowIdx, processIdx),
		process.ValidateImage(workflowIdx, processIdx),
		process.ValidateReplicas(workflowIdx, processIdx),
		process.ValidateGPU(workflowIdx, processIdx),
		process.ValidateConfig(workflowIdx, processIdx),
		process.ValidateObjectStore(workflowIdx, processIdx),
		process.ValidateSecrets(workflowIdx, processIdx),
		process.ValidateSubscriptions(workflowIdx, processIdx),
		process.ValidateNetworking(workflowIdx, processIdx),
		process.ValidateResourceLimits(workflowIdx, processIdx),
	)
}

func (process *Process) ValidateName(workflowIdx, processIdx int) error {
	return validateName(
		process.Name,
		fmt.Sprintf("krt.workflows[%d].processes[%d].name", workflowIdx, processIdx),
	)
}

func (process *Process) ValidateType(workflowIdx, processIdx int) error {
	if !process.Type.IsValid() {
		return errors.InvalidProcessTypeError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].type", workflowIdx, processIdx),
		)
	}

	return nil
}

func (process *Process) ValidateImage(workflowIdx, processIdx int) error {
	if process.Image == "" {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].image", workflowIdx, processIdx),
		)
	}

	return nil
}

func (process *Process) ValidateReplicas(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) ValidateGPU(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) ValidateConfig(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) ValidateObjectStore(workflowIdx, processIdx int) error {
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

func (process *Process) ValidateSecrets(workflowIdx, processIdx int) error {
	return nil
}

func (process *Process) ValidateSubscriptions(workflowIdx, processIdx int) error {
	if process.Type == ProcessTypeTrigger {
		return nil
	}

	if process.Subscriptions == nil || len(process.Subscriptions) == 0 {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].subscriptions",
				workflowIdx,
				processIdx,
			),
		)
	}

	return nil
}

// validateSubscritpionRelationships checks if subscriptions for all processes are valid
// inisde a workflow context.
//
// All requirements for subscritpions to be valid can be found in the readme.
func validateSubscritpionRelationships(processes []Process, workflowIdx int) error {
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
						fmt.Sprintf(subscriptionLocation, workflowIdx, processIdx, subscription),
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
			cleanSubscription := strings.Split(subscription, ".")[0]

			if process.Name == cleanSubscription {
				totalError = errors.Join(totalError, errors.CannotSubscribeToItselfError(
					fmt.Sprintf(subscriptionLocation, workflowIdx, processIdx, subscription),
				))

				continue
			}

			subscribedProcessType, processExists := processTypesByNames[cleanSubscription]
			if !processExists {
				totalError = errors.Join(totalError, errors.CannotSubscribeToNonExistentProcessError(
					subscription,
					fmt.Sprintf("krt.workflows[%d].processes[%d]", workflowIdx, processIdx),
				))

				continue
			}

			if !isValidSubscription(process.Type, subscribedProcessType) {
				totalError = errors.Join(totalError, errors.InvalidProcessSubscriptionError(
					string(process.Type),
					string(subscribedProcessType),
					fmt.Sprintf(subscriptionLocation, workflowIdx, processIdx, subscription),
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

func (process *Process) ValidateNetworking(workflowIdx, processIdx int) error {
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

func (process *Process) ValidateResourceLimits(workflowIdx, processIdx int) error {
	if process.ResourceLimits == nil {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits", workflowIdx, processIdx),
		)
	}

	return errors.Join(
		process.ValidateCPU(workflowIdx, processIdx),
		process.ValidateMemory(workflowIdx, processIdx),
	)
}

func (process *Process) ValidateCPU(workflowIdx, processIdx int) error {
	if process.ResourceLimits.CPU == nil {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.CPU", workflowIdx, processIdx),
		)
	}

	if process.ResourceLimits.CPU.Request == "" {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.CPU.request", workflowIdx, processIdx),
		)
	}

	var (
		totalError  error
		requestForm cpuForm
		limitForm   cpuForm
	)

	requestOk, requestForm := isValidCPU(process.ResourceLimits.CPU.Request)
	if !requestOk {
		totalError = errors.Join(
			totalError,
			errors.InvalidProcessCPUError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.CPU.request", workflowIdx, processIdx),
			),
		)
	}

	if process.ResourceLimits.CPU.Limit != "" {
		var limitOk bool
		limitOk, limitForm = isValidCPU(process.ResourceLimits.CPU.Limit)

		if !limitOk {
			totalError = errors.Join(
				totalError,
				errors.InvalidProcessCPUError(
					fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.CPU.limit", workflowIdx, processIdx),
				),
			)
		}
	} else {
		process.ResourceLimits.CPU.Limit = process.ResourceLimits.CPU.Request
		limitForm = requestForm
	}

	if totalError == nil {
		totalError = compareRequestLimitCPU(
			process.ResourceLimits.CPU.Request, process.ResourceLimits.CPU.Limit, requestForm, limitForm, workflowIdx, processIdx,
		)
	}

	return totalError
}

func (process *Process) ValidateMemory(workflowIdx, processIdx int) error {
	if process.ResourceLimits.Memory == nil {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.memory", workflowIdx, processIdx),
		)
	}

	if process.ResourceLimits.Memory.Request == "" {
		return errors.MissingRequiredFieldError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.memory.request", workflowIdx, processIdx),
		)
	}

	var totalError error

	requestOk := isValidMemory(process.ResourceLimits.Memory.Request)
	if !requestOk {
		totalError = errors.Join(
			totalError,
			errors.InvalidProcessMemoryError(
				fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.memory.request", workflowIdx, processIdx),
			),
		)
	}

	if process.ResourceLimits.Memory.Limit != "" {
		limitOk := isValidMemory(process.ResourceLimits.Memory.Limit)
		if !limitOk {
			totalError = errors.Join(
				totalError,
				errors.InvalidProcessMemoryError(
					fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.memory.limit", workflowIdx, processIdx),
				),
			)
		}
	} else {
		process.ResourceLimits.Memory.Limit = process.ResourceLimits.Memory.Request
	}

	if totalError == nil {
		totalError = compareRequestLimitMemory(
			process.ResourceLimits.Memory.Request, process.ResourceLimits.Memory.Limit, workflowIdx, processIdx,
		)
	}

	return totalError
}
