package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_returns_false_when_no_matching_step_is_defined(t *testing.T) {
	testee := &gourd_steps{}

	_, match, _ := testee.matching_step("undefined step")
	assert.False(t, match)
}

func Test_matching_a_step_returns_a_non_empty_id(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "pattern"
	testee.add_step(pattern)

	id, match, _ := testee.matching_step(pattern)
	assert.True(t, match)
	assert.NotEqual(t, id, "")
}

func Test_two_different_steps_return_a_different_id(t *testing.T) {
	testee := &gourd_steps{}

	first_pattern := "a pattern"
	second_pattern := "another pattern"
	testee.add_step(first_pattern)
	testee.add_step(second_pattern)

	first_id, match, _ := testee.matching_step(first_pattern)
	assert.True(t, match)

	second_id, match, _ := testee.matching_step(second_pattern)
	assert.True(t, match)

	assert.NotEqual(t, first_id, second_id)
}

func Test_begin_scenario_creates_new_context_by_calling_the_injected_function(t *testing.T) {
	called := false
	testee := &gourd_steps{}
	testee.new_context = func() interface{} {
		called = true
		return nil
	}

	testee.begin_scenario()

	assert.True(t, called)
}

func Test_begin_scenario_ignores_the_injected_function_when_it_is_nil(t *testing.T) {
	testee := &gourd_steps{}

	assert.NotPanics(t, func() {
		testee.begin_scenario()
	})
}

func Test_step_pattern_is_matched_as_a_regular_expression(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "match number (\\d+)"
	testee.add_step(pattern)

	_, match, _ := testee.matching_step("match number 42")
	assert.True(t, match)
}

func Test_a_capturing_group_is_returned_with_value_and_position(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "match number (\\d+)"
	testee.add_step(pattern)

	_, _, arguments := testee.matching_step("match number 42")

	expected := []capturing_group{
		capturing_group{value: "42", position: 13}}

	assert.Equal(t, expected, arguments)
}
