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

func (parser *wire_protocol_parser) parse(command []byte) string {
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

func (parser *wire_protocol_parser) evaluate(request string, command []json.RawMessage) string {
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

func (parser *wire_protocol_parser) step_matches(parameters json.RawMessage) string {
	var patterns map[string]string
	json.Unmarshal(parameters, &patterns)
	pattern := patterns["name_to_match"]
	id, matches, arguments := parser.steps.matching_step(pattern)
	if matches {
		return success_response_steps(id, arguments)
	}
	return success_response_no_match()
}

func (parser *wire_protocol_parser) snippet_text(parameters json.RawMessage) string {
	var snippet map[string]string
	json.Unmarshal(parameters, &snippet)
	return success_response_snippet(snippet["step_name"], snippet["step_keyword"])
}

func (parser *wire_protocol_parser) invoke(parameters json.RawMessage) string {
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

func (parser *wire_protocol_parser) begin_scenario() string {
	parser.steps.begin_scenario()
	return success_response()
}

func (parser *wire_protocol_parser) end_scenario() string {
	return success_response()
}

func wire_response(status string) []interface{} {
	return []interface{}{status}
}

func wire_success() []interface{} {
	return wire_response("success")
}

func wire_pending() []interface{} {
	return wire_response("pending")
}

func wire_fail() []interface{} {
	return wire_response("fail")
}

func encode_json(response []interface{}) string {
	encoded, _ := json.Marshal(response)
	return string(encoded)
}

func success_response() string {
	return encode_json(wire_success())
}

func success_response_steps(id string, arguments []capturing_group) string {

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
	return encode_json(response)
}

func success_response_no_match() string {
	response := append(wire_success(), []interface{}{})
	return encode_json(response)
}

func success_response_snippet(name string, keyword string) string {
	escaped_name := strings.Replace(name, "\"", "\\\"", -1)
	snippet_text := `cucumber.` + keyword + `("` + escaped_name + `").Pending()`

	response := append(wire_success(), snippet_text)

	return encode_json(response)
}

func pending_response() string {
	return encode_json(wire_pending())
}

func fail_response(message string) string {
	response := append(wire_fail(), map[string]string{"message": message})
	return encode_json(response)
}
