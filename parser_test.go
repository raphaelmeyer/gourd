package gourd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_parser_returns_success_to_begin_scenario(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	steps.On("begin_scenario").Return().Once()

	command := []byte(`["begin_scenario"]`)
	response := testee.parse(command)

	assert.Equal(t, `["success"]`, response)
}

func Test_parser_notifies_steps_when_a_new_scenario_begins(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	steps.On("begin_scenario").Return().Once()

	command := []byte(`["begin_scenario"]`)
	testee.parse(command)

	steps.Mock.AssertExpectations(t)
}

func Test_parser_returns_success_to_end_scenario(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["end_scenario"]`)
	response := testee.parse(command)

	assert.Equal(t, `["success"]`, response)
}

func Test_parser_asks_for_matching_step_with_given_pattern(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	pattern := "Given pattern"
	steps.On("matching_step", pattern).Return("", false, []argument{}).Once()

	command := []byte(`["step_matches",{"name_to_match":"` + pattern + `"}]`)
	testee.parse(command)

	steps.Mock.AssertExpectations(t)
}

func Test_parser_returns_success_and_empty_array_for_undefined_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	steps.On("matching_step", mock.Anything).Return("", false, []argument{}).Once()

	command := []byte(`["step_matches",{"name_to_match":"undefined step"}]`)
	response := testee.parse(command)

	assert.Equal(t, `["success",[]]`, response)
}

func Test_parser_returns_success_and_id_for_defined_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "123"
	pattern := "defined step"
	steps.On("matching_step", pattern).Return(id, true, []argument{}).Once()

	command := []byte(`["step_matches",{"name_to_match":"` + pattern + `"}]`)
	response := testee.parse(command)

	expected_response := `["success",[{"id":"` + id + `","args":[]}]]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_failure_for_unknown_command(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["unknown_command"]`)
	response := testee.parse(command)

	expected_response := `["fail",{"message":"unknown command: unknown_command"}]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_snippet_text_for_given(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["snippet_text",{"step_keyword":"Given","multiline_arg_class":"","step_name":"Step"}]`)
	response := testee.parse(command)

	expected_response := `["success","cucumber.Given(\"Step\").Pending()"]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_snippet_text_for_when(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["snippet_text",{"step_keyword":"When","multiline_arg_class":"","step_name":"when step"}]`)
	response := testee.parse(command)

	expected_response := `["success","cucumber.When(\"when step\").Pending()"]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_escapes_special_characters_in_snippet(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`["snippet_text",{"step_keyword":"Then","multiline_arg_class":"","step_name":"step with '\"quotes\"'"}]`)
	response := testee.parse(command)

	expected_response := `["success","cucumber.Then(\"step with '\\\"quotes\\\"'\").Pending()"]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_invokes_a_step_with_the_given_id(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	id := "7"
	steps.On("invoke_step", id, mock.Anything).Return(success, "").Once()

	command := []byte(`["invoke",{"id":"` + id + `","args":[]}]`)
	testee.parse(command)

	steps.Mock.AssertExpectations(t)
}

func Test_parser_returns_pending_when_invoking_a_pending_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	steps.On("invoke_step", mock.Anything, mock.Anything).Return(pending, "").Once()

	command := []byte(`["invoke",{"id":"13","args":[]}]`)
	response := testee.parse(command)

	expected_response := `["pending"]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_success_when_invoking_a_passing_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	steps.On("invoke_step", mock.Anything, mock.Anything).Return(success, "").Once()

	command := []byte(`["invoke",{"id":"24","args":[]}]`)
	response := testee.parse(command)

	expected_response := `["success"]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_fail_when_invoking_a_failing_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	steps.On("invoke_step", mock.Anything, mock.Anything).Return(fail, "").Once()

	command := []byte(`["invoke",{"id":"35","args":[]}]`)
	response := testee.parse(command)

	expected_response := `["fail",{"message":""}]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_failure_message_of_failing_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	message := "failure message"
	steps.On("invoke_step", mock.Anything, mock.Anything).Return(fail, message).Once()

	command := []byte(`["invoke",{"id":"46","args":[]}]`)
	response := testee.parse(command)

	expected_response := `["fail",{"message":"` + message + `"}]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_fail_when_the_command_is_malformed_json(t *testing.T) {
	testee := &wire_protocol_parser{}

	command := []byte(`[}"this is not valid json"`)
	response := testee.parse(command)

	expected_response := `["fail",{"message":"invalid command"}]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_a_capturing_group_as_an_argument(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	arguments := []argument{
		argument{value: "value", position: 5}}

	steps.On("matching_step", mock.Anything).Return("47", true, arguments).Once()

	command := []byte(`["step_matches",{"name_to_match":"some value"}]`)
	response := testee.parse(command)

	expected_response := `["success",[{"id":"47","args":[{"val":"value","pos":5}]}]]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_returns_all_capturing_groups_as_arguments(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	arguments := []argument{
		argument{value: "some", position: 0},
		argument{value: "value", position: 5},
		argument{value: "match", position: 14}}

	steps.On("matching_step", mock.Anything).Return("47", true, arguments).Once()

	command := []byte(`["step_matches",{"name_to_match":"some value to match"}]`)
	response := testee.parse(command)

	expected_response := `["success",[{"id":"47","args":[{"val":"some","pos":0},{"val":"value","pos":5},{"val":"match","pos":14}]}]]`
	assert.Equal(t, expected_response, response)
}

func Test_parser_passes_arguments_to_invoked_step(t *testing.T) {
	steps := &steps_mock{}
	testee := &wire_protocol_parser{steps}

	expected_arguments := []string{"value"}
	steps.On("invoke_step", mock.Anything, expected_arguments).Return(success, "").Once()

	command := []byte(`["invoke",{"id":"123","args":["value"]}]`)
	testee.parse(command)

	steps.Mock.AssertExpectations(t)
}
