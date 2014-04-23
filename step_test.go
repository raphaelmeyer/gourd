package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_a_step_is_pending_by_default(t *testing.T) {
	testee := &gourd_step{"pattern"}

	assert.True(t, testee.is_pending())
}

// Test_a_step_marked_as_passing_is_not_pending
// Test_a_step_with_an_action_is_not_pending
