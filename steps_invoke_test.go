package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_invoking_a_pending_step_returns_pending(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Pending()
	id, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, pending)
}

func Test_a_step_is_pending_by_default(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	testee.add_step(pattern)
	id, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, pending)
}

func Test_invoking_a_step_that_is_set_to_always_pass_returns_success(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Pass()
	id, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, success)
}

func Test_invoking_a_step_that_is_set_to_always_fail_returns_fail(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Fail()
	id, _ := testee.matching_step(pattern)

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
	id, _ := testee.matching_step(pattern)

	testee.invoke_step(id)

	assert.True(t, executed)
}

func Test_invoking_a_step_whos_action_does_not_panic_returns_success(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}) {
	})
	id, _ := testee.matching_step(pattern)

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
	id, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id)

	assert.Equal(t, result, fail)
}

func Test_invoking_a_failing_step_returns_the_failure_message(t *testing.T) {
	t.Log("pending")
}
