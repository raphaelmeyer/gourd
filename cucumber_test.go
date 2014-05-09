package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_statement_given_adds_and_returns_a_new_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &gourd_cucumber{steps, nil, DefaultPort}

	pattern := "arbitrary step pattern"

	steps.On("add_step", pattern).Return(&gourd_step{}).Once()

	step := testee.Given(pattern)

	assert.NotNil(t, step)
	steps.Mock.AssertExpectations(t)
}

func Test_statement_when_adds_and_returns_a_new_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &gourd_cucumber{steps, nil, DefaultPort}

	pattern := "arbitrary step pattern"

	steps.On("add_step", pattern).Return(&gourd_step{}).Once()

	step := testee.When(pattern)

	assert.NotNil(t, step)
	steps.Mock.AssertExpectations(t)
}

func Test_statement_then_adds_and_returns_a_new_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &gourd_cucumber{steps, nil, DefaultPort}

	pattern := "arbitrary step pattern"

	steps.On("add_step", pattern).Return(&gourd_step{}).Once()

	step := testee.Then(pattern)

	assert.NotNil(t, step)
	steps.Mock.AssertExpectations(t)
}

func Test_run_starts_the_wire_server_on_the_default_port(t *testing.T) {
	server := &server_mock{}
	testee := &gourd_cucumber{nil, server, DefaultPort}

	server.On("listen", DefaultPort).Return().Once()

	testee.Run()

	server.Mock.AssertExpectations(t)
}

func Test_run_starts_the_wire_server_on_the_specified_port(t *testing.T) {
	server := &server_mock{}
	testee := &gourd_cucumber{nil, server, DefaultPort}

	var specified_port uint = 2345
	server.On("listen", specified_port).Return().Once()

	testee.SetPort(specified_port)
	testee.Run()

	server.Mock.AssertExpectations(t)
}
