package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_returns_false_when_no_matching_step_is_defined(t *testing.T) {
	testee := &gourd_steps{}

	_, match := testee.matching_step("undefined step")
	assert.False(t, match)
}

func Test_matching_a_step_returns_a_non_empty_id(t *testing.T) {
	testee := &gourd_steps{}

	pattern := "pattern"
	testee.add_step(pattern)

	id, match := testee.matching_step(pattern)
	assert.True(t, match)
	assert.NotEqual(t, id, "")
}

func Test_two_different_steps_return_a_different_id(t *testing.T) {
	testee := &gourd_steps{}

	first_pattern := "a pattern"
	second_pattern := "another pattern"
	testee.add_step(first_pattern)
	testee.add_step(second_pattern)

	first_id, match := testee.matching_step(first_pattern)
	assert.True(t, match)

	second_id, match := testee.matching_step(second_pattern)
	assert.True(t, match)

	assert.NotEqual(t, first_id, second_id)
}
