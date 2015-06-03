package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_command_parser_parses_begin_scenario_command(t *testing.T) {
	testee := wire_command_parser{}

	command_string := []byte(`["begin_scenario"]`)

	command := testee.parse(command_string)

	expected := &wire_command_begin_scenario{}
	assert.IsType(t, expected, command)
}

func Test_command_parser_parses_end_scenario_command(t *testing.T) {
	testee := wire_command_parser{}

	command_string := []byte(`["end_scenario"]`)

	command := testee.parse(command_string)

	expected := &wire_command_end_scenario{}
	assert.IsType(t, expected, command)
}
