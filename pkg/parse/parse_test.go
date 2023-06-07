//go:build unit

package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectKrtFile(t *testing.T) {
	krt, err := ParseFile("./files/correct_krt.yaml")
	assert.NoError(t, err)

	err = krt.Validate()
	assert.NoError(t, err)
}
