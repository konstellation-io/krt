package krt

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

func isValidWorkflowType(workflowType string) bool {
	var workflowTypeMap = map[string]WorkflowType{
		string(WorkflowTypeData):     WorkflowTypeData,
		string(WorkflowTypeTraining): WorkflowTypeTraining,
		string(WorkflowTypeFeedback): WorkflowTypeFeedback,
		string(WorkflowTypeServing):  WorkflowTypeServing,
	}

	_, ok := workflowTypeMap[workflowType]

	return ok
}

type Process struct {
	Name          string              `yaml:"name"`
	Type          ProcessType         `yaml:"type"`
	Image         string              `yaml:"image"`
	Replicas      *int                `yaml:"replicas" default:"1"`
	GPU           *bool               `yaml:"gpu" default:"false" `
	Config        map[string]string   `yaml:"config"`
	ObjectStore   *ProcessObjectStore `yaml:"objectStore"`
	Secrets       map[string]string   `yaml:"secrets"`
	Subscriptions []string            `yaml:"subscriptions"`
	Networking    *ProcessNetworking  `yaml:"networking"`
}

type ProcessType string

const (
	ProcessTypeTrigger ProcessType = "trigger"
	ProcessTypeTask    ProcessType = "task"
	ProcessTypeExit    ProcessType = "exit"
)

func isValidProcessType(processType string) bool {
	var processTypeMap = map[string]ProcessType{
		string(ProcessTypeTrigger): ProcessTypeTrigger,
		string(ProcessTypeTask):    ProcessTypeTask,
		string(ProcessTypeExit):    ProcessTypeExit,
	}

	_, ok := processTypeMap[processType]

	return ok
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

func isValidObjectStoreScope(scope string) bool {
	var objectStoreScopeMap = map[string]ObjectStoreScope{
		string(ObjectStoreScopeProduct):  ObjectStoreScopeProduct,
		string(ObjectStoreScopeWorkflow): ObjectStoreScopeWorkflow,
	}

	_, ok := objectStoreScopeMap[scope]

	return ok
}

type ProcessNetworking struct {
	TargetPort          int                `yaml:"targetPort" default:"9000" `
	TargetProtocol      NetworkingProtocol `yaml:"targetProtocol" default:"TCP" `
	DestinationPort     int                `yaml:"destinationPort" default:"9000" `
	DestinationProtocol NetworkingProtocol `yaml:"destinationProtocol" default:"TCP" `
}

type NetworkingProtocol string

const (
	NetworkingProtocolTCP NetworkingProtocol = "TCP"
	NetworkingProtocolUDP NetworkingProtocol = "UDP"
)

func isValidNetworkingProtocol(protocol string) bool {
	var networkingProtocolMap = map[string]NetworkingProtocol{
		string(NetworkingProtocolTCP): NetworkingProtocolTCP,
		string(NetworkingProtocolUDP): NetworkingProtocolUDP,
	}

	_, ok := networkingProtocolMap[protocol]

	return ok
}
