package gourd

import (
	"github.com/stretchr/testify/mock"
)

type steps_mock struct {
	mock.Mock
}

func (steps *steps_mock) matching_step(step string) (bool, string) {
	args := steps.Mock.Called(step)
	return args.Bool(0), args.String(1)
}

func (steps *steps_mock) add_step(pattern string) Step {
	args := steps.Mock.Called(pattern)
	return args.Get(0).(*gourd_step)
}

func (steps *steps_mock) invoke_step(id string) step_result {
	args := steps.Mock.Called(id)
	return args.Get(0).(step_result)
}
