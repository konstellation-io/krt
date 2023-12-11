package parse

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
)

// ParseYamlToKrt parses a Krt struct from a given yaml bytes.
func ParseYamlToKrt(krtYaml []byte) (*krt.Krt, error) {
	// talk about this shadow import
	var parsedKrt krt.Krt

	err := yaml.Unmarshal(krtYaml, &parsedKrt)
	if err != nil {
		return nil, errors.InvalidYamlError(err)
	}

	defaults.MustSet(&parsedKrt)

	return &parsedKrt, nil
}

// ParseFileToKrt parses a Krt struct from a given filename.
//
// File must be in yaml format.
func ParseFileToKrt(yamlFile string) (*krt.Krt, error) {
	krtYml, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, errors.ReadingFileError(err)
	}

	return ParseYamlToKrt(krtYml)
}

// ParseKrtToYaml parses a Krt struct to yaml bytes.
func ParseKrtToYaml(krtStruct *krt.Krt) ([]byte, error) {
	return yaml.Marshal(krtStruct)
}
