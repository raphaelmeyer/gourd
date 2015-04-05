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
	id, _, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id, []string{})

	assert.Equal(t, result, pending)
}

func Test_a_step_is_pending_by_default(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	testee.add_step(pattern)
	id, _, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id, []string{})

	assert.Equal(t, result, pending)
}

func Test_invoking_a_step_that_is_set_to_always_pass_returns_success(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Pass()
	id, _, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id, []string{})

	assert.Equal(t, result, success)
}

func Test_invoking_a_step_that_is_set_to_always_fail_returns_fail(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Fail()
	id, _, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id, []string{})

	assert.Equal(t, result, fail)
}

func Test_invoking_a_step_with_an_unknown_id_fails(t *testing.T) {
	testee := &gourd_steps{}

	result, _ := testee.invoke_step("unknown id", []string{})

	assert.Equal(t, result, fail)
}

func Test_invoking_a_step_executes_the_defined_action(t *testing.T) {
	testee := &gourd_steps{}

	executed := false
	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}, arguments Arguments) {
		executed = true
	})
	id, _, _ := testee.matching_step(pattern)

	testee.invoke_step(id, []string{})

	assert.True(t, executed)
}

func Test_invoking_a_step_whos_action_does_not_panic_returns_success(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}, arguments Arguments) {
	})
	id, _, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id, []string{})

	assert.Equal(t, result, success)
}

func Test_invoking_a_step_whos_action_panics_returns_fail(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}, arguments Arguments) {
		panic("")
	})
	id, _, _ := testee.matching_step(pattern)

	result, _ := testee.invoke_step(id, []string{})

	assert.Equal(t, result, fail)
}

func Test_invoking_a_failing_step_returns_the_failure_message(t *testing.T) {
	testee := &gourd_steps{}

	expected_message := "error message"
	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}, arguments Arguments) {
		panic(expected_message)
	})
	id, _, _ := testee.matching_step(pattern)

	_, message := testee.invoke_step(id, []string{})

	assert.Equal(t, expected_message, message)
}

func Test_invoke_step_passes_the_context_created_in_begin_scenario(t *testing.T) {
	type context_type struct {
		value int
	}
	expected_context := &context_type{123}

	testee := &gourd_steps{}
	testee.new_context = func() interface{} {
		return expected_context
	}

	var actual_context *context_type
	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Do(func(context interface{}, arguments Arguments) {
		actual_context = context.(*context_type)
	})
	id, _, _ := testee.matching_step(pattern)

	testee.begin_scenario()
	testee.invoke_step(id, []string{})

	assert.Equal(t, expected_context, actual_context)
	assert.Exactly(t, expected_context, actual_context)
}

func Test_invoke_step_passes_string_arguments_to_the_action(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	id, _, _ := testee.matching_step(pattern)

	args := []string{"some string", "another one"}

	step.Do(func(context interface{}, arguments Arguments) {
		assert.Equal(t, "some string", arguments.String(0))
		assert.Equal(t, "another one", arguments.String(1))
	})

	testee.invoke_step(id, args)
}

func Test_invoking_a_step_fails_when_the_step_tries_to_access_an_invalid_argument_index(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	id, _, _ := testee.matching_step(pattern)

	step.Do(func(context interface{}, arguments Arguments) {
		arguments.String(0)
	})

	result, _ := testee.invoke_step(id, []string{})

	assert.Equal(t, fail, result)
}
