package main

type Krt struct {
	Name        string
	Description string
	Version     string
	Config      map[string]string
	Workflows   []Workflow
}

type Workflow struct {
	Name      string
	Type      WorkflowType
	Config    map[string]string
	Processes []Process
}

type WorkflowType string

const (
	WorkflowTypeData     WorkflowType = "data"
	WorkflowTypeTraining WorkflowType = "training"
	WorkflowTypeFeedback WorkflowType = "feedback"
	WorkflowTypeServing  WorkflowType = "serving"
)

type Process struct {
	Name          string
	Type          ProcessType
	Build         ProcessBuild
	Replicas      int
	GPU           bool
	Config        map[string]string
	ObjectStore   ProcessObjectStore
	Secrets       []string
	Subscriptions []string
	Networking    ProcessNetworking
}

type ProcessType string

const (
	ProcessTypeTrigger ProcessType = "trigger"
	ProcessTypeTask    ProcessType = "task"
	ProcessTypeExit    ProcessType = "exit"
)

type ProcessBuild struct {
	Image      string
	Dockerfile string
}

type ProcessObjectStore struct {
	Name  string
	Scope ObjectStoreScope
}

type ObjectStoreScope string

const (
	ObjectStoreScopeProduct  ObjectStoreScope = "product"
	ObjectStoreScopeWorkflow ObjectStoreScope = "workflow"
)

type ProcessNetworking struct {
	TargetPort          int
	TargetProtocol      NetworkingProtocol
	DestinationPort     int
	DestinationProtocol NetworkingProtocol
}

type NetworkingProtocol string

const (
	NetworkingProtocolTCP NetworkingProtocol = "TCP"
	NetworkingProtocolUDP NetworkingProtocol = "UDP"
)
