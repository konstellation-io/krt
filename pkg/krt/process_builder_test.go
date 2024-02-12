//go:build unit

package krt_test

import "github.com/konstellation-io/krt/pkg/krt"

type ProcessBuilder struct {
	process *krt.Process
}

func NewProcessBuilder() *ProcessBuilder {
	return &ProcessBuilder{
		process: &krt.Process{
			Name:          "test-process",
			Type:          krt.ProcessTypeTask,
			Image:         "test-image",
			Replicas:      nil,
			GPU:           nil,
			Config:        nil,
			ObjectStore:   nil,
			Secrets:       nil,
			Subscriptions: []string{"test-trigger"},
			Networking:    nil,
			ResourceLimits: &krt.ProcessResourceLimits{
				CPU: &krt.ResourceLimit{
					Request: "100m",
					Limit:   "200m",
				},
				Memory: &krt.ResourceLimit{
					Request: "100M",
					Limit:   "200M",
				},
			},
			NodeSelectors: nil,
		},
	}
}

func (pb *ProcessBuilder) WithNodeSelectors(nodeSelectors map[string]string) *ProcessBuilder {
	pb.process.NodeSelectors = nodeSelectors
	return pb
}

func (pb *ProcessBuilder) Build() *krt.Process {
	return pb.process
}
