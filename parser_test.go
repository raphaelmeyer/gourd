package gourd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_parser_returns_success_to_begin_scenario(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["begin_scenario"]` + "\n")
	response := testee.parse(command)

	assert.Equal(t, response, `["success"]`+"\n")
}

func Test_parser_returns_success_to_end_scenario(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte("[\"end_scenario\"]\n")
	response := testee.parse(command)

	assert.Equal(t, response, "[\"success\"]\n")
}

func Test_parser_asks_for_matching_step_with_given_pattern(t *testing.T) {
	steps := &stepsMock{}
	testee := &wire_protocol_parser{steps}

	pattern := "Given pattern"
	steps.On("matching_step", pattern).Return(false, 0).Once()

	command := []byte(`["step_matches",{"name_to_match":"` + pattern + `"}]` + "\n")
	_ = testee.parse(command)
}

func Test_parser_returns_success_and_empty_array_for_undefined_step(t *testing.T) {
	steps := &stepsMock{}
	testee := &wire_protocol_parser{steps}

	steps.On("matching_step", mock.Anything).Return(false, 0).Once()

	command := []byte(`["step_matches",{"name_to_match":"undefined step"}]` + "\n")
	response := testee.parse(command)

	assert.Equal(t, response, `["success",[]]`+"\n")
}

func Test_parser_returns_success_and_id_for_defined_step(t *testing.T) {
	steps := &stepsMock{}
	testee := &wire_protocol_parser{steps}

	id := 1
	pattern := "defined step"
	steps.On("matching_step", pattern).Return(true, id).Once()

	command := []byte(`["step_matches",{"name_to_match":"` + pattern + `"}]` + "\n")
	response := testee.parse(command)

	expected_response := fmt.Sprintf(`["success",[{"id":"%d", "args":[]}]]`+"\n", id)

	assert.Equal(t, response, expected_response)

	steps.Mock.AssertExpectations(t)
}

func Test_parser_returns_the_id_of_the_matching_step(t *testing.T) {
	steps := &stepsMock{}
	testee := &wire_protocol_parser{steps}

	id := 5
	pattern := "defined step"
	steps.On("matching_step", pattern).Return(true, id).Once()

	command := []byte(`["step_matches",{"name_to_match":"` + pattern + `"}]` + "\n")
	response := testee.parse(command)

	expected_response := fmt.Sprintf(`["success",[{"id":"%d", "args":[]}]]`+"\n", id)

	assert.Equal(t, response, expected_response)

	steps.Mock.AssertExpectations(t)
}

func Test_parser_returns_failure_for_unknown_command(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["unknown_command"]` + "\n")
	response := testee.parse(command)

	expected_response := `["fail",{"message":"unknown command"}]` + "\n"
	assert.Equal(t, response, expected_response)
}

func Test_parser_returns_snippet_text_for_given(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["snippet_text",{"step_keyword":"Given","multiline_arg_class":"","step_name":"Step"}]` + "\n")
	response := testee.parse(command)

	expected_response := `["success","cucumber.Given(\"Step\").Pending()\n"]` + "\n"
	assert.Equal(t, response, expected_response)
}

func Test_parser_returns_snippet_text_for_when(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["snippet_text",{"step_keyword":"When","multiline_arg_class":"","step_name":"when step"}]` + "\n")
	response := testee.parse(command)

	expected_response := `["success","cucumber.When(\"when step\").Pending()\n"]` + "\n"
	assert.Equal(t, response, expected_response)
}
