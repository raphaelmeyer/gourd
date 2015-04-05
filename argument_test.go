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

func Test_access_string_arguments(t *testing.T) {
	testee := &gourd_arguments{[]string{"some string", "another one"}}

	assert.Equal(t, "some string", testee.String(0))
	assert.Equal(t, "another one", testee.String(1))
}

func Test_access_integer_arguemnts(t *testing.T) {
	testee := &gourd_arguments{[]string{"123", "-45"}}

	assert.Equal(t, 123, testee.Int(0))
	assert.Equal(t, -45, testee.Int(1))
}

func Test_arguments_can_be_accessed_multiple_times(t *testing.T) {
	testee := &gourd_arguments{[]string{"some string", "-71"}}

	assert.Equal(t, -71, testee.Int(1))
	assert.Equal(t, -71, testee.Int(1))
}

func Test_access_order_does_not_matter(t *testing.T) {
	input := []string{"something", "6273"}
	testee := &gourd_arguments{input}

	assert.Equal(t, 6273, testee.Int(1))
	assert.Equal(t, "something", testee.String(0))
}

func Test_panics_if_argument_cannot_be_converted_into_an_integer(t *testing.T) {
	testee := &gourd_arguments{[]string{"not a number"}}

	assert.Panics(t, func() {
		testee.Int(0)
	})
}
