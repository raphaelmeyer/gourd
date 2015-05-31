package gourd

import (
	"encoding/json"
	"strings"
)

type parser interface {
	parse(command []byte) string
}

type wire_protocol_parser struct {
	steps steps
}

func (cmd *generic_wire_command) execute() wire_response {
	return &generic_wire_response{}
}

func (response *generic_wire_response) encode() string {
	return ""
}

func (parser *wire_protocol_parser) parse(command []byte) string {

	cmd := parse_wire_command(command)
	response := cmd.execute()
	_ = response.encode()

	raw_response := parser.parse_command(command)
	return encode_json(raw_response)
}

func (parser *wire_protocol_parser) parse_command(command []byte) []interface{} {
	var raw_command []json.RawMessage
	if err := json.Unmarshal(command, &raw_command); err != nil {
		return fail_response("invalid command")
	}

	var request string
	if err := json.Unmarshal(raw_command[0], &request); err != nil {
		return fail_response("invalid command")
	}

	return parser.evaluate(request, raw_command)
}

func (parser *wire_protocol_parser) evaluate(request string, command []json.RawMessage) []interface{} {
	switch request {
	case "step_matches":
		return parser.step_matches(command[1])
	case "begin_scenario":
		return parser.begin_scenario()
	case "end_scenario":
		return parser.end_scenario()
	case "snippet_text":
		return parser.snippet_text(command[1])
	case "invoke":
		return parser.invoke(command[1])
	}
	return fail_response("unknown command: " + request)
}

func (parser *wire_protocol_parser) step_matches(parameters json.RawMessage) []interface{} {
	var patterns map[string]string
	json.Unmarshal(parameters, &patterns)
	pattern := patterns["name_to_match"]
	id, matches, arguments := parser.steps.matching_step(pattern)
	if matches {
		return success_response_steps(id, arguments)
	}
	return success_response_no_match()
}

func (parser *wire_protocol_parser) snippet_text(parameters json.RawMessage) []interface{} {
	var snippet map[string]string
	json.Unmarshal(parameters, &snippet)
	return success_response_snippet(snippet["step_name"], snippet["step_keyword"])
}

func (parser *wire_protocol_parser) invoke(parameters json.RawMessage) []interface{} {
	var invoke map[string]json.RawMessage
	json.Unmarshal(parameters, &invoke)
	var id string
	json.Unmarshal(invoke["id"], &id)
	var arguments []string
	json.Unmarshal(invoke["args"], &arguments)
	result, message := parser.steps.invoke_step(id, arguments)
	switch result {
	case fail:
		return fail_response(message)
	case pending:
		return pending_response()
	}
	return success_response()
}

func (parser *wire_protocol_parser) begin_scenario() []interface{} {
	parser.steps.begin_scenario()
	return success_response()
}

func (parser *wire_protocol_parser) end_scenario() []interface{} {
	return success_response()
}

func wire_response_(status string) []interface{} {
	return []interface{}{status}
}

func wire_success() []interface{} {
	return wire_response_("success")
}

func wire_pending() []interface{} {
	return wire_response_("pending")
}

func wire_fail() []interface{} {
	return wire_response_("fail")
}

func encode_json(response []interface{}) string {
	encoded, _ := json.Marshal(response)
	return string(encoded)
}

func success_response() []interface{} {
	return wire_success()
}

func success_response_steps(id string, arguments []capturing_group) []interface{} {

	type wire_argument struct {
		Value     string `json:"val"`
		Positions uint   `json:"pos"`
	}

	type wire_match struct {
		Id   string          `json:"id"`
		Args []wire_argument `json:"args"`
	}

	args := []wire_argument{}
	for _, argument := range arguments {
		args = append(args, wire_argument{argument.value, argument.position})
	}

	match := wire_match{id, args}
	response := append(wire_success(), []wire_match{match})

	return response
}

func success_response_no_match() []interface{} {
	response := append(wire_success(), []interface{}{})
	return response
}

func success_response_snippet(name string, keyword string) []interface{} {
	escaped_name := strings.Replace(name, "\"", "\\\"", -1)
	snippet_text := `cucumber.` + keyword + `("` + escaped_name + `").Pending()`

	response := append(wire_success(), snippet_text)

	return response
}

func pending_response() []interface{} {
	return wire_pending()
}

func fail_response(message string) []interface{} {
	response := append(wire_fail(), map[string]string{"message": message})
	return response
}
