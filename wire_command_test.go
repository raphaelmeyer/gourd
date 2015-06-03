package gourd

import (
	//"github.com/stretchr/testify/assert"
	"testing"
)

func Test_begin_scneario_notifies_the_steps(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_command_begin_scenario{}

	steps.On("begin_scenario").Return().Once()

	testee.execute(steps)

	steps.Mock.AssertExpectations(t)
}
