package gourd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func Test_parser_returns_the_id_of_a_defined_step(t *testing.T) {
	steps := &StepManagerMock{}
	testee := &CommandParser{steps}

	command := `["step_matches",{"name_to_match":"defined step"}]` + "\n"
	response := testee.Parse(command)

	assert.Equal(t, response, `["success",[{"id":"1", "args":[]}]]` + "\n")
}

type StepManagerMock struct {
	mock.Mock
}

func (step_manager *StepManagerMock) MatchingStep(pattern string) (bool, int) {
	args := step_manager.Mock.Called(pattern)
	return args.Bool(0), args.Int(1)
}

