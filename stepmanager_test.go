package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_returns_false_when_no_matching_step_is_defined(t *testing.T) {
	testee := &StepManager{}

	found, _ := testee.MatchingStep("undefined step")
	assert.False(t, found)
}

func Test_matching_a_step_returns_a_positive_id(t *testing.T) {
	testee := &StepManager{}

	pattern := "pattern"
	testee.AddStep(pattern)

	found, id := testee.MatchingStep(pattern)
	assert.True(t, found)
	assert.True(t, 0 < id)
}
