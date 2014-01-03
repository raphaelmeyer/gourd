package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parser_returns_success_to_begin_scenario(t *testing.T) {
	testee := &CommandParser{}

	command := "[\"begin_scenario\"]\n"
	response := testee.Parse(command)

	assert.Equal(t, response, "[\"success\"]\n")
}

func Test_parser_returns_success_to_end_scenario(t *testing.T) {
	testee := &CommandParser{}

	command := "[\"end_scenario\"]\n"
	response := testee.Parse(command)

	assert.Equal(t, response, "[\"success\"]\n")
}

func Test_parser_returns_success_and_empty_array_for_undefined_step(t *testing.T) {
	testee := &CommandParser{}

	command := `["step_matches",{"name_to_match":"undefined step"}]` + "\n"
	response := testee.Parse(command)

	assert.Equal(t, response, `["success",[]]` + "\n")
}

