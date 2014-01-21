package gourd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"fmt"
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

func TODO_parser_extracts_the_step_pattern(t *testing.T) {
	assert.Fail(t, "Pending")
}

func Test_parser_returns_success_and_empty_array_for_undefined_step(t *testing.T) {
	steps := &StepManagerMock{}
	testee := &CommandParser{steps}

	steps.On("MatchingStep", mock.Anything).Return(false, 0).Once()

	command := `["step_matches",{"name_to_match":"undefined step"}]` + "\n"
	response := testee.Parse(command)

	assert.Equal(t, response, `["success",[]]` + "\n")
}

func Test_parser_returns_the_id_of_a_defined_step(t *testing.T) {
	steps := &StepManagerMock{}
	testee := &CommandParser{steps}

	id := 1
	pattern := "defined step"
	steps.On("MatchingStep", pattern).Return(true, id).Once()

	command := `["step_matches",{"name_to_match":"` + pattern + `"}]` + "\n"
	response := testee.Parse(command)

	expected_response := fmt.Sprintf(`["success",[{"id":"%d", "args":[]}]]` + "\n", id)

	assert.Equal(t, response, expected_response)

	steps.Mock.AssertExpectations(t)
}

type StepManagerMock struct {
	mock.Mock
}

func (steps *StepManagerMock) MatchingStep(pattern string) (bool, int) {
	args := steps.Mock.Called(pattern)
	return args.Bool(0), args.Int(1)
}

