package gourd

import (
	"encoding/json"
	"fmt"
	"strings"
)

type parser interface {
	parse(command []byte) string
}

type wire_protocol_parser struct {
	steps steps
}

func (parser *wire_protocol_parser) parse(command []byte) string {
	var data []json.RawMessage
	if err := json.Unmarshal(command, &data); err != nil {
		return fail_response("invalid command")
	}

	var request string
	if err := json.Unmarshal(data[0], &request); err != nil {
		return fail_response("invalid command")
	}

	return parser.evaluate(request, data)
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
	return `["success",[]]`
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
		return `["pending"]`
	}
	return `["success"]`
}

func (parser *wire_protocol_parser) begin_scenario() string {
	parser.steps.begin_scenario()
	return `["success"]`
}

func (parser *wire_protocol_parser) end_scenario() string {
	return `["success"]`
}

func build_arguments_string(arguments []capturing_group) string {
	arguments_string := ""
	for i, argument := range arguments {
		if i > 0 {
			arguments_string += ","
		}
		position := fmt.Sprintf("%d", argument.position)
		arguments_string += `{"val":"` + argument.value + `","pos":` + position + `}`
	}
	return arguments_string
}

func fail_response(message string) string {
	return `["fail",{"message":"` + message + `"}]`
}

func success_response_steps(id string, arguments []capturing_group) string {
	arguments_string := build_arguments_string(arguments)
	return `["success",[{"id":"` + id + `","args":[` + arguments_string + `]}]]`
}

func success_response_snippet(name string, keyword string) string {
	escaped_name := strings.Replace(name, "\"", "\\\"", -1)
	snippet_text, _ := json.Marshal(`cucumber.` + keyword + `("` + escaped_name + `").Pending()`)
	return `["success",` + string(snippet_text) + `]`
}
