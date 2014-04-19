package gourd

import (
	"github.com/stretchr/testify/mock"
)

type stepsMock struct {
	mock.Mock
}

func (steps *stepsMock) matching_step(step string) (bool, int) {
	args := steps.Mock.Called(step)
	return args.Bool(0), args.Int(1)
}

func (steps *stepsMock) add_step(pattern string) Step {
	args := steps.Mock.Called(pattern)
	return args.Get(0).(*gourd_step)
}
