package gourd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_success_response_encodes_to_json_string(t *testing.T) {
	testee := &wire_response_success{}

	result := testee.encode()

	assert.Equal(t, `["success"]`, result)
}
