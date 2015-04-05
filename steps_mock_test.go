package gourd

import (
	"github.com/stretchr/testify/mock"
)

type steps_mock struct {
	mock.Mock
}

func (steps *steps_mock) begin_scenario() {
	_ = steps.Mock.Called()
}

func (steps *steps_mock) matching_step(step string) (string, bool, []argument) {
	args := steps.Mock.Called(step)
	return args.String(0), args.Bool(1), args.Get(2).([]argument)
}

func (steps *steps_mock) add_step(pattern string) Step {
	args := steps.Mock.Called(pattern)
	return args.Get(0).(*gourd_step)
}

func (steps *steps_mock) invoke_step(id string, arguments []string) (step_result, string) {
	args := steps.Mock.Called(id, arguments)
	return args.Get(0).(step_result), args.String(1)
}
