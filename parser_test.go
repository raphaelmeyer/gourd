package gourd

import (
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
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	pattern := "Given pattern"
	steps.On("matching_step", pattern).Return("", false).Once()

	command := []byte(`["step_matches",{"name_to_match":"` + pattern + `"}]` + "\n")
	_ = testee.parse(command)
}

func Test_parser_returns_success_and_empty_array_for_undefined_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	steps.On("matching_step", mock.Anything).Return("", false).Once()

	command := []byte(`["step_matches",{"name_to_match":"undefined step"}]` + "\n")
	response := testee.parse(command)

	assert.Equal(t, response, `["success",[]]`+"\n")
}

func Test_parser_returns_success_and_id_for_defined_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "123"
	pattern := "defined step"
	steps.On("matching_step", pattern).Return(id, true).Once()

	command := []byte(`["step_matches",{"name_to_match":"` + pattern + `"}]` + "\n")
	response := testee.parse(command)

	expected_response := `["success",[{"id":"` + id + `", "args":[]}]]` + "\n"

	assert.Equal(t, response, expected_response)

	steps.Mock.AssertExpectations(t)
}

func Test_parser_returns_failure_for_unknown_command(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["unknown_command"]` + "\n")
	response := testee.parse(command)

	expected_response := `["fail",{"message":"unknown command: unknown_command"}]` + "\n"
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

func Test_parser_invokes_a_step_with_the_given_id(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "7"
	steps.On("invoke_step", id).Return(success, "").Once()

	command := []byte(`["invoke",{"id":"` + id + `","args":[]}]` + "\n")
	testee.parse(command)

	steps.Mock.AssertExpectations(t)
}

func Test_parser_returns_pending_when_invoking_a_pending_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "13"
	steps.On("invoke_step", id).Return(pending, "").Once()

	command := []byte(`["invoke",{"id":"` + id + `","args":[]}]` + "\n")
	response := testee.parse(command)

	expected_response := `["pending"]` + "\n"
	assert.Equal(t, response, expected_response)
}

func Test_parser_returns_success_when_invoking_a_passing_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "37"
	steps.On("invoke_step", id).Return(success, "").Once()

	command := []byte(`["invoke",{"id":"` + id + `","args":[]}]` + "\n")
	response := testee.parse(command)

	expected_response := `["success"]` + "\n"
	assert.Equal(t, response, expected_response)
}

func Test_parser_returns_fail_when_invoking_a_failing_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "123"
	steps.On("invoke_step", id).Return(fail, "").Once()

	command := []byte(`["invoke",{"id":"` + id + `","args":[]}]` + "\n")
	response := testee.parse(command)

	expected_response := `["fail",{"message":""}]` + "\n"
	assert.Equal(t, response, expected_response)
}

func Test_parser_returns_failure_message_of_failing_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "4"
	message := "failure message"
	steps.On("invoke_step", id).Return(fail, message).Once()

	command := []byte(`["invoke",{"id":"` + id + `","args":[]}]` + "\n")
	response := testee.parse(command)

	expected_response := `["fail",{"message":"` + message + `"}]` + "\n"
	assert.Equal(t, response, expected_response)
}
