package parse

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"

	"github.com/konstellation-io/krt/pkg/errors"
	"github.com/konstellation-io/krt/pkg/krt"
)

// ParseKrt parses a Krt struct from a given yaml bytes.
func ParseKrt(krtYaml []byte) (*krt.Krt, error) {
	// talk about this shadow import
	var parsedKrt krt.Krt

	err := yaml.Unmarshal(krtYaml, &parsedKrt)
	if err != nil {
		return nil, errors.InvalidYamlError(err)
	}

	err = setDefaultsForStruct(&parsedKrt)
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

func setDefaultsForStruct(v interface{}) error {
	err := defaults.Set(v)
	if err != nil {
		return errors.SetDefaultsError(err)
	}

	return nil
}
