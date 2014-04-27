package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_invoking_a_step_that_is_set_to_always_pass_returns_success(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Pass()
	_, id := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, success)
}

func Test_invoking_a_step_that_is_set_to_always_fail_returns_fail(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Fail()
	_, id := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, fail)
}

func Test_invoking_a_step_with_an_unknown_id_fails(t *testing.T) {
	testee := &gourd_steps{}

	result, _ := testee.invoke_step("unknown id")

	assert.Equal(t, result, fail)
}

func Test_invoking_a_step_executes_the_defined_action(t *testing.T) {
	testee := &gourd_steps{}

	executed := false
	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}) {
		executed = true
	})
	_, id := testee.matching_step(pattern)

	testee.invoke_step(id)

	assert.True(t, executed)
}

func Test_invoking_a_step_whos_action_does_not_panic_returns_success(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}) {
	})
	_, id := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, success)
}

func Test_invoking_a_step_whos_action_panics_returns_fail(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}) {
		panic("")
	})
	_, id := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, fail)
}

func Test_invoking_a_failing_step_returns_the_failure_message(t *testing.T) {
	t.Log("pending")
}
