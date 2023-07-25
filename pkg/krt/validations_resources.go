package krt

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/units"
	"github.com/konstellation-io/krt/pkg/errors"
)

type cpuForm int

const (
	cpuFormFractional cpuForm = iota
	cpuFormMilli
)

func isValidCPU(cpu string) (bool, cpuForm) {
	fractionalCPU := regexp.MustCompile(`^\d\.\d$`)
	milliCPU := regexp.MustCompile(`^\d{3}m$`)

	if fractionalCPU.MatchString(cpu) {
		return true, cpuFormFractional
	}

	if milliCPU.MatchString(cpu) {
		return true, cpuFormMilli
	}

	return false, 0
}

func getCPUValue(cpu string, form cpuForm) (float64, error) {
	if form == cpuFormFractional {
		cpuValue, err := strconv.ParseFloat(cpu, 32)
		if err != nil {
			return 0, err
		}

		return cpuValue * 1000, nil
	} else {
		cpuValue, err := strconv.ParseFloat(strings.ReplaceAll(cpu, "m", ""), 32)
		if err != nil {
			return 0, err
		}

		return cpuValue, nil
	}
}

func compareRequestLimitCPU(request, limit string, requestForm, limitForm cpuForm, workflowIdx, processIdx int) error {
	requestValue, err := getCPUValue(request, requestForm)
	if err != nil {
		return err
	}

	limitValue, err := getCPUValue(limit, limitForm)
	if err != nil {
		return err
	}

	if limitValue < requestValue {
		return errors.InvalidProcessCPURelationError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].CPU", workflowIdx, processIdx),
		)
	}

	return nil
}

func isValidMemory(memory string) bool {
	megaBMemory := regexp.MustCompile(`^\d+(Ei?|Pi?|Ti?|Gi?|Mi?|k|Ki)$`)
	return megaBMemory.MatchString(memory)
}

func getMemoryValue(memory string) (int64, error) {
	return units.ParseStrictBytes(memory + "B") // Added B to match expected format
}

func compareRequestLimitMemory(request, limit string, workflowIdx, processIdx int) error {
	requestValue, err := getMemoryValue(request)
	if err != nil {
		return err
	}

	limitValue, err := getMemoryValue(limit)
	if err != nil {
		return err
	}

	if limitValue < requestValue {
		return errors.InvalidProcessMemoryRelationError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].memory", workflowIdx, processIdx),
		)
	}

	return nil
}
