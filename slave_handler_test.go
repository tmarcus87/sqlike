package sqlike

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestSlaveSelectionHandler_UnmarshalText(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name string
	}{
		{name: "random"},
		{name: "round_robbin"},
	}

	for _, t := range tests {
		var h SlaveSelectionHandler

		err := h.UnmarshalText([]byte(t.name))
		asserts.Nil(err)
	}
}

func TestSlaveSelectionHandler_UnmarshalYAML(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name string
	}{
		{name: "random"},
		{name: "round_robbin"},
	}

	for _, t := range tests {
		var h SlaveSelectionHandler
		err := yaml.Unmarshal([]byte(t.name), &h)
		asserts.Nil(err)
	}
}