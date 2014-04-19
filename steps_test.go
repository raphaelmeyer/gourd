package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_returns_false_when_no_matching_step_is_defined(t *testing.T) {
	testee := &cucumberSteps{}

	found, _ := testee.matching_step("undefined step")
	assert.False(t, found)
}

func Test_matching_a_step_returns_a_positive_id(t *testing.T) {
	testee := &cucumberSteps{}

	pattern := "pattern"
	testee.add_step(pattern)

	found, id := testee.matching_step(pattern)
	assert.True(t, found)
	assert.True(t, 0 < id)
}

func Test_two_different_steps_return_a_different_id(t *testing.T) {
	testee := &cucumberSteps{}

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
