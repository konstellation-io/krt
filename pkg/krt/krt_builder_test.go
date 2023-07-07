//go:build unit

package krt_test

import "github.com/konstellation-io/krt/pkg/krt"

type KrtBuilder struct {
	krtYaml *krt.Krt
}

func NewKrtBuilder() *KrtBuilder {
	return &KrtBuilder{
		krtYaml: &krt.Krt{
			Version:     "v1.0.0",
			Description: "Test description",
			Workflows: []krt.Workflow{
				{
					Name: "test-workflow",
					Type: krt.WorkflowTypeTraining,
					Processes: []krt.Process{
						{
							Name:  "test-trigger",
							Type:  krt.ProcessTypeTrigger,
							Image: "test-trigger-image",
							Subscriptions: []string{
								"test-exit",
							},
						},
						{
							Name:  "test-exit",
							Type:  krt.ProcessTypeExit,
							Image: "test-exit-image",
							Subscriptions: []string{
								"test-trigger",
							},
						},
					},
				},
			},
		},
	}
}

func (k *KrtBuilder) WithVersion(version string) *KrtBuilder {
	k.krtYaml.Version = version
	return k
}

func (k *KrtBuilder) WithDescription(description string) *KrtBuilder {
	k.krtYaml.Description = description
	return k
}

func (k *KrtBuilder) WithVersionConfig(config map[string]string) *KrtBuilder {
	k.krtYaml.Config = config
	return k
}

func (k *KrtBuilder) WithWorkflows(workflows []krt.Workflow) *KrtBuilder {
	k.krtYaml.Workflows = workflows
	return k
}

func (k *KrtBuilder) WithWorkflowName(name string) *KrtBuilder {
	k.krtYaml.Workflows[0].Name = name
	return k
}

func (k *KrtBuilder) WithWorkflowType(workflowType krt.WorkflowType) *KrtBuilder {
	k.krtYaml.Workflows[0].Type = workflowType
	return k
}

func (k *KrtBuilder) WithProcesses(processes []krt.Process) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes = processes
	return k
}

func (k *KrtBuilder) WithProcessName(name string, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Name = name
	return k
}

func (k *KrtBuilder) WithProcessType(processType krt.ProcessType, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Type = processType
	return k
}

func (k *KrtBuilder) WithProcessImage(image string, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Image = image
	return k
}

func (k *KrtBuilder) WithProcessReplicas(replicas *int, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Replicas = replicas
	return k
}

func (k *KrtBuilder) WithProcessGPU(gpu *bool, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].GPU = gpu
	return k
}

func (k *KrtBuilder) WithProcessConfig(config map[string]string, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Config = config
	return k
}

func (k *KrtBuilder) WithProcessObjectStore(objectStore *krt.ProcessObjectStore, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].ObjectStore = objectStore
	return k
}

func (k *KrtBuilder) WithProcessSecrets(secrets map[string]string, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Secrets = secrets
	return k
}

func (k *KrtBuilder) WithProcessSubscriptions(subscriptions []string, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Subscriptions = subscriptions
	return k
}

func (k *KrtBuilder) WithProcessNetworking(networking *krt.ProcessNetworking, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Networking = networking
	return k
}

func (k *KrtBuilder) Build() *krt.Krt {
	return k.krtYaml
}
