package main

type Krt struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Version     string            `yaml:"version"`
	Config      map[string]string `yaml:"config"`
	Workflows   []Workflow        `yaml:"workflows"`
}

type Workflow struct {
	Name      string            `yaml:"name"`
	Type      WorkflowType      `yaml:"type"`
	Config    map[string]string `yaml:"config"`
	Processes []Process         `yaml:"processes"`
}

type WorkflowType string

const (
	WorkflowTypeData     WorkflowType = "data"
	WorkflowTypeTraining WorkflowType = "training"
	WorkflowTypeFeedback WorkflowType = "feedback"
	WorkflowTypeServing  WorkflowType = "serving"
)

var WorkflowTypeMap = map[string]WorkflowType{
	"data":     WorkflowTypeData,
	"training": WorkflowTypeTraining,
	"feedback": WorkflowTypeFeedback,
	"serving":  WorkflowTypeServing,
}

type Process struct {
	Name          string             `yaml:"name"`
	Type          ProcessType        `yaml:"type"`
	Build         ProcessBuild       `yaml:"build"`
	Replicas      int                `yaml:"replicas"`
	GPU           bool               `yaml:"gpu"`
	Config        map[string]string  `yaml:"config"`
	ObjectStore   ProcessObjectStore `yaml:"objectStore"`
	Secrets       []string           `yaml:"secrets"`
	Subscriptions []string           `yaml:"subscriptions"`
	Networking    ProcessNetworking  `yaml:"networking"`
}

type ProcessType string

const (
	ProcessTypeTrigger ProcessType = "trigger"
	ProcessTypeTask    ProcessType = "task"
	ProcessTypeExit    ProcessType = "exit"
)

var ProcessTypeMap = map[string]ProcessType{
	"trigger": ProcessTypeTrigger,
	"task":    ProcessTypeTask,
	"exit":    ProcessTypeExit,
}

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

var ObjectStoreScopeMap = map[string]ObjectStoreScope{
	"product":  ObjectStoreScopeProduct,
	"workflow": ObjectStoreScopeWorkflow,
}

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

var NetworkingProtocolMap = map[string]NetworkingProtocol{
	"TCP": NetworkingProtocolTCP,
	"UDP": NetworkingProtocolUDP,
}
