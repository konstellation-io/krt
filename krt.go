package main

type Krt struct {
	Name        string            `yaml:"name" validate:"required"`
	Description string            `yaml:"description" validate:"required"`
	Version     string            `yaml:"version" validate:"required,resource-name,lt=20"`
	Config      map[string]string `yaml:"config" validate:"omitempty"`
	Workflows   []Workflow        `yaml:"workflows" validate:"required,dive"`
}

type Workflow struct {
	Name      string            `yaml:"name" validate:"required,resource-name,lt=20"`
	Type      WorkflowType      `yaml:"type" validate:"required"`
	Config    map[string]string `yaml:"config" validate:"omitempty"`
	Processes []Process         `yaml:"processes" validate:"required,dive"`
}

type WorkflowType string

const (
	WorkflowTypeData     WorkflowType = "data"
	WorkflowTypeTraining WorkflowType = "training"
	WorkflowTypeFeedback WorkflowType = "feedback"
	WorkflowTypeServing  WorkflowType = "serving"
)

type Process struct {
	Name          string             `yaml:"name" validate:"required,resource-name,lt=20"`
	Type          ProcessType        `yaml:"type" validate:"required"`
	Build         ProcessBuild       `yaml:"build" validate:"required"`
	Replicas      int                `yaml:"replicas" validate:"omitempty"`
	GPU           bool               `yaml:"gpu" validate:"omitempty"`
	Config        map[string]string  `yaml:"config" validate:"omitempty"`
	ObjectStore   ProcessObjectStore `yaml:"objectStore" validate:"omitempty"`
	Secrets       []string           `yaml:"secrets" validate:"omitempty"`
	Subscriptions []string           `yaml:"subscriptions" validate:"required"`
	Networking    ProcessNetworking  `yaml:"networking" validate:"omitempty"`
}

type ProcessType string

const (
	ProcessTypeTrigger ProcessType = "trigger"
	ProcessTypeTask    ProcessType = "task"
	ProcessTypeExit    ProcessType = "exit"
)

type ProcessBuild struct {
	Image      string `yaml:"image"`
	Dockerfile string `yaml:"dockerfile"`
}

type ProcessObjectStore struct {
	Name  string           `yaml:"name"`
	Scope ObjectStoreScope `yaml:"scope"`
}

type ObjectStoreScope string

const (
	ObjectStoreScopeProduct  ObjectStoreScope = "product"
	ObjectStoreScopeWorkflow ObjectStoreScope = "workflow"
)

type ProcessNetworking struct {
	TargetPort          int                `yaml:"targetPort"`
	TargetProtocol      NetworkingProtocol `yaml:"targetProtocol"`
	DestinationPort     int                `yaml:"destinationPort"`
	DestinationProtocol NetworkingProtocol `yaml:"destinationProtocol"`
}

type NetworkingProtocol string

const (
	NetworkingProtocolTCP NetworkingProtocol = "TCP"
	NetworkingProtocolUDP NetworkingProtocol = "UDP"
)
