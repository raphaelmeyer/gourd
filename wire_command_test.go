package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_begin_scneario_notifies_the_steps(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_command_begin_scenario{}

	steps.On("begin_scenario").Return().Once()

	testee.execute(steps)

	steps.Mock.AssertExpectations(t)
}

func Test_begin_scenario_returns_success_response(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_command_begin_scenario{}

	steps.On("begin_scenario").Return().Once()

	response := testee.execute(steps)

	expected := &wire_response_success{}
	assert.IsType(t, expected, response)
}

func Test_end_scenario_returns_success_response(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_command_end_scenario{}

	response := testee.execute(steps)

	expected := &wire_response_success{}
	assert.IsType(t, expected, response)
}
