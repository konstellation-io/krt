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

func getCPUValue(cpu string, form cpuForm) float64 {
	if form == cpuFormFractional {
		cpuValue, _ := strconv.ParseFloat(cpu, 32)

		return cpuValue * 1000
	} else {
		cpuValue, _ := strconv.ParseFloat(strings.ReplaceAll(cpu, "m", ""), 32)

		return cpuValue
	}
}

func compareRequestLimitCPU(request, limit string, requestForm, limitForm cpuForm, workflowIdx, processIdx int) error {
	requestValue := getCPUValue(request, requestForm)

	limitValue := getCPUValue(limit, limitForm)

	if limitValue < requestValue {
		return errors.InvalidProcessCPURelationError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.CPU", workflowIdx, processIdx),
		)
	}

	return nil
}

func isValidMemory(memory string) bool {
	megaBMemory := regexp.MustCompile(`^\d+(Ei?|Pi?|Ti?|Gi?|Mi?|k|Ki)$`)
	return megaBMemory.MatchString(memory)
}

func getMemoryValue(memory string) int64 {
	strictBytes, _ := units.ParseStrictBytes(memory + "B") // Added B to match expected format
	return strictBytes
}

func compareRequestLimitMemory(request, limit string, workflowIdx, processIdx int) error {
	requestValue := getMemoryValue(request)

	limitValue := getMemoryValue(limit)

	if limitValue < requestValue {
		return errors.InvalidProcessMemoryRelationError(
			fmt.Sprintf("krt.workflows[%d].processes[%d].resourceLimits.memory", workflowIdx, processIdx),
		)
	}

	return nil
}
