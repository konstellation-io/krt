//go:build unit

package main

type KrtBuilder struct {
	krtYaml *Krt
}

func NewKrtBuilder() *KrtBuilder {
	return &KrtBuilder{
		krtYaml: &Krt{
			Name:        "Test KRT",
			Description: "Test description",
			Version:     "version-name",
			Workflows: []Workflow{
				{
					Name: "test-workflow",
					Type: WorkflowTypeTraining,
					Processes: []Process{
						{
							Name:  "test-trigger",
							Type:  ProcessTypeTrigger,
							Build: ProcessBuild{Image: "test-trigger-image"},
							Subscriptions: []string{
								"test-exit",
							},
						},
						{
							Name:  "test-exit",
							Type:  ProcessTypeExit,
							Build: ProcessBuild{Image: "test-exit-image"},
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

func (k *KrtBuilder) WithProcessesForWorkflow(processes []Process, workflowIdx int) *KrtBuilder {
	k.krtYaml.Workflows[workflowIdx].Processes = processes
	return k
}

func (k *KrtBuilder) Build() *Krt {
	return k.krtYaml
}
