package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_panics_when_accessed_without_being_initialized_with_string_slice(t *testing.T) {
	testee := &gourd_arguments{}

	assert.Panics(t, func() {
		testee.String(0)
	})
}

func Test_panics_when_accessed_index_is_out_range(t *testing.T) {
	testee := &gourd_arguments{[]string{}}

	assert.Panics(t, func() {
		testee.String(0)
	})
}
