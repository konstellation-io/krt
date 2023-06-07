package parse

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
)

const (
	defaultReplicas = 1
	defaultGPU      = false
	defaultPort     = 9000
	defaultProtocol = "TCP"
)

// ParseKrt parses a Krt struct from a given yaml bytes.
func ParseKrt(krtYaml []byte) (*krt.Krt, error) {
	var parsedKrt krt.Krt

	err := yaml.Unmarshal(krtYaml, &parsedKrt)
	if err != nil {
		return nil, errors.InvalidYamlError(err)
	}

	setDefaults(&parsedKrt)

	return &parsedKrt, nil
}

// ParseFile parses a Krt struct from a given filename.
func ParseFile(yamlFile string) (*krt.Krt, error) {
	krtYml, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, errors.ReadingFileError(err)
	}

	return ParseKrt(krtYml)
}

// set a logger maybe?
// talk about this shadow import
func setDefaults(krtStruct *krt.Krt) {
	for _, workflow := range krtStruct.Workflows {
		for _, process := range workflow.Processes {
			setDefaultReplicas(process.Replicas)
			setDefaultGPU(process.GPU)
			setDefaultNetworkingProtocol(process.Networking)
		}
	}
}

func setDefaultReplicas(replicas *int) {
	if replicas == nil || *replicas < 1 {
		*replicas = defaultReplicas
	}
}

func setDefaultGPU(gpu *bool) {
	if gpu == nil {
		*gpu = defaultGPU
	}
}

func setDefaultNetworkingProtocol(networking *krt.ProcessNetworking) {
	if networking == nil {
		return
	}

	if networking.TargetPort == 0 {
		networking.TargetPort = defaultPort
	}

	if networking.TargetProtocol == "" {
		networking.TargetProtocol = defaultProtocol
	}

	if networking.DestinationPort == 0 {
		networking.DestinationPort = defaultPort
	}

	if networking.DestinationProtocol == "" {
		networking.DestinationProtocol = defaultProtocol
	}
}
