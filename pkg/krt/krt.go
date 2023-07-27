package krt

type Krt struct {
	Version     string            `yaml:"version"`
	Description string            `yaml:"description"`
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

func (wt WorkflowType) IsValid() bool {
	var workflowTypeMap = map[string]WorkflowType{
		string(WorkflowTypeData):     WorkflowTypeData,
		string(WorkflowTypeTraining): WorkflowTypeTraining,
		string(WorkflowTypeFeedback): WorkflowTypeFeedback,
		string(WorkflowTypeServing):  WorkflowTypeServing,
	}

	_, ok := workflowTypeMap[string(wt)]

	return ok
}

const (
	DefaultNumberOfReplicas = 1
	DefaultGPUValue         = false
)

type Process struct {
	Name           string                 `yaml:"name"`
	Type           ProcessType            `yaml:"type"`
	Image          string                 `yaml:"image"`
	Replicas       *int                   `yaml:"replicas" default:"1"`
	GPU            *bool                  `yaml:"gpu" default:"false" `
	Config         map[string]string      `yaml:"config"`
	ObjectStore    *ProcessObjectStore    `yaml:"objectStore"`
	Secrets        map[string]string      `yaml:"secrets"`
	Subscriptions  []string               `yaml:"subscriptions"`
	Networking     *ProcessNetworking     `yaml:"networking"`
	ResourceLimits *ProcessResourceLimits `yaml:"resourceLimits"`
}

type ProcessType string

const (
	ProcessTypeTrigger ProcessType = "trigger"
	ProcessTypeTask    ProcessType = "task"
	ProcessTypeExit    ProcessType = "exit"
)

func (pt ProcessType) IsValid() bool {
	var processTypeMap = map[string]ProcessType{
		string(ProcessTypeTrigger): ProcessTypeTrigger,
		string(ProcessTypeTask):    ProcessTypeTask,
		string(ProcessTypeExit):    ProcessTypeExit,
	}

	_, ok := processTypeMap[string(pt)]

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

func (s ObjectStoreScope) IsValid() bool {
	var objectStoreScopeMap = map[string]ObjectStoreScope{
		string(ObjectStoreScopeProduct):  ObjectStoreScopeProduct,
		string(ObjectStoreScopeWorkflow): ObjectStoreScopeWorkflow,
	}

	_, ok := objectStoreScopeMap[string(s)]

	return ok
}

const (
	DefaultProtocol = NetworkingProtocolTCP
)

type ProcessNetworking struct {
	TargetPort      int                `yaml:"targetPort"`
	DestinationPort int                `yaml:"destinationPort"`
	Protocol        NetworkingProtocol `yaml:"protocol" default:"TCP" `
}

type NetworkingProtocol string

const (
	NetworkingProtocolTCP NetworkingProtocol = "TCP"
	NetworkingProtocolUDP NetworkingProtocol = "UDP"
)

func (np NetworkingProtocol) IsValid() bool {
	var networkingProtocolMap = map[string]NetworkingProtocol{
		string(NetworkingProtocolTCP): NetworkingProtocolTCP,
		string(NetworkingProtocolUDP): NetworkingProtocolUDP,
	}

	_, ok := networkingProtocolMap[string(np)]

	return ok
}

type ProcessCPU struct {
	Request string `yaml:"request"`
	Limit   string `yaml:"limit"`
}

type ProcessMemory struct {
	Request string `yaml:"request"`
	Limit   string `yaml:"limit"`
}

type ProcessResourceLimits struct {
	CPU    *ProcessCPU    `yaml:"CPU"`
	Memory *ProcessMemory `yaml:"memory"`
}
