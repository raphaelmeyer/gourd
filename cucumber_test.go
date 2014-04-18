package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_a_given_statement_adds_a_new_step(t *testing.T) {
	steps := &stepsMock{}
	testee := &gourdCucumber{steps, nil}

	pattern := "arbitrary step pattern"

	steps.On("addStep", pattern).Return(&Step{}).Once()

	step := testee.Given(pattern)

	assert.NotNil(t, step)
	steps.Mock.AssertExpectations(t)
}
