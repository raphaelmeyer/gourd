package gourd

import (
	"github.com/stretchr/testify/mock"
)

type stepsMock struct {
	mock.Mock
}

func (steps *stepsMock) matchingStep(step string) (bool, int) {
	args := steps.Mock.Called(step)
	return args.Bool(0), args.Int(1)
}

func (steps *stepsMock) addStep(pattern string) {
	steps.Mock.Called(pattern)
}
