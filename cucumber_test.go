package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_statement_given_adds_and_returns_a_new_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &gourd_cucumber{steps, nil}

	pattern := "arbitrary step pattern"

	steps.On("add_step", pattern).Return(&gourd_step{}).Once()

	step := testee.Given(pattern)

	assert.NotNil(t, step)
	steps.Mock.AssertExpectations(t)
}

func Test_run_starts_the_wire_server(t *testing.T) {
	server := &server_mock{}
	testee := &gourd_cucumber{nil, server}

	server.On("listen").Return()

	testee.Run()

	server.Mock.AssertExpectations(t)
}
