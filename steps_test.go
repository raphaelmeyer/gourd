package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_returns_false_when_no_matching_step_is_defined(t *testing.T) {
	testee := &gourd_steps{}

	found, _ := testee.matching_step("undefined step")
	assert.False(t, found)
}

func Test_matching_a_step_returns_a_non_empty_id(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "pattern"
	testee.add_step(pattern)

	found, id := testee.matching_step(pattern)
	assert.True(t, found)
	assert.NotEqual(t, id, "")
}

func Test_two_different_steps_return_a_different_id(t *testing.T) {
	testee := &gourd_steps{}

	first_pattern := "a pattern"
	second_pattern := "another pattern"
	testee.add_step(first_pattern)
	testee.add_step(second_pattern)

	found, first_id := testee.matching_step(first_pattern)
	assert.True(t, found)

	found, second_id := testee.matching_step(second_pattern)
	assert.True(t, found)

	assert.NotEqual(t, first_id, second_id)
}

func Test_invoking_a_pending_step_returns_pending(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	step := testee.add_step(pattern)
	step.Pending()
	_, id := testee.matching_step(pattern)

	result := testee.invoke_step(id)

	assert.Equal(t, result, pending)
}

func Test_a_step_is_pending_by_default(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "arbitrary step pattern"
	testee.add_step(pattern)
	_, id := testee.matching_step(pattern)

	result := testee.invoke_step(id)

	assert.Equal(t, result, pending)
}
