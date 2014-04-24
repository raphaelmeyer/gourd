package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_a_step_is_pending_by_default(t *testing.T) {
	testee := new_step("pattern")

	assert.True(t, testee.is_pending())
}

func Test_a_step_can_explicitly_be_marked_pending(t *testing.T) {
	testee := new_step("pattern")
	testee.Pending()

	assert.True(t, testee.is_pending())
}

func Test_a_step_marked_to_always_pass_is_not_pending(t *testing.T) {
	testee := new_step("pattern")
	testee.Pass()

	assert.False(t, testee.is_pending())
}
