package parse

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
	// talk about this shadow import
)

// ParseKrt parses a Krt struct from a given yaml bytes.
func ParseKrt(krtYaml []byte) (*krt.Krt, error) {
	var parsedKrt krt.Krt

	err := yaml.Unmarshal(krtYaml, &parsedKrt)
	if err != nil {
		return nil, errors.InvalidYamlError(err)
	}

	err = defaults.Set(&parsedKrt)
	if err != nil {
		return nil, errors.SetDefaultsError(err)
	}

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
