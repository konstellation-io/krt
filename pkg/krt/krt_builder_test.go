//go:build unit

package krt

type KrtBuilder struct {
	krtYaml *Krt
}

func NewKrtBuilder() *KrtBuilder {
	return &KrtBuilder{
		krtYaml: &Krt{
			Name:        "test-krt",
			Description: "Test description",
			Version:     "v1.0",
			Workflows: []Workflow{
				{
					Name: "test-workflow",
					Type: WorkflowTypeTraining,
					Processes: []Process{
						{
							Name:  "test-trigger",
							Type:  ProcessTypeTrigger,
							Image: "test-trigger-image",
							Subscriptions: []string{
								"test-exit",
							},
						},
						{
							Name:  "test-exit",
							Type:  ProcessTypeExit,
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

func (k *KrtBuilder) WithName(name string) *KrtBuilder {
	k.krtYaml.Name = name
	return k
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

func (k *KrtBuilder) WithWorkflows(workflows []Workflow) *KrtBuilder {
	k.krtYaml.Workflows = workflows
	return k
}

func (k *KrtBuilder) WithWorkflowName(name string) *KrtBuilder {
	k.krtYaml.Workflows[0].Name = name
	return k
}

func (k *KrtBuilder) WithWorkflowType(workflowType WorkflowType) *KrtBuilder {
	k.krtYaml.Workflows[0].Type = workflowType
	return k
}

func (k *KrtBuilder) WithProcesses(processes []Process) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes = processes
	return k
}

func (k *KrtBuilder) WithProcessName(name string, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Name = name
	return k
}

func (k *KrtBuilder) WithProcessType(processType ProcessType, processIdx int) *KrtBuilder {
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

func (k *KrtBuilder) WithProcessObjectStore(objectStore *ProcessObjectStore, processIdx int) *KrtBuilder {
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

func (k *KrtBuilder) WithProcessNetworking(networking *ProcessNetworking, processIdx int) *KrtBuilder {
	k.krtYaml.Workflows[0].Processes[processIdx].Networking = networking
	return k
}

func (k *KrtBuilder) Build() *Krt {
	return k.krtYaml
}
