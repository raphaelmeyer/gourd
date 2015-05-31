package gourd

import (
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
	"testing"
)

func Test_(t *testing.T) {
	command_string := []byte(`["begin_scenario"]`)
	expected := &wire_command_begin_scenario{}

	command := parse_wire_command(command_string)

	assert.Equal(t, expected, command)
}

