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
		return `["fail",{"message":"invalid command"}]`
	}

	var request string
	if err := json.Unmarshal(data[0], &request); err != nil {
		return `["fail",{"message":"invalid command"}]`
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
		return `["success"]`
	case "snippet_text":
		return parser.snippet_text(command[1])
	case "invoke":
		return parser.invoke(command[1])
	}
	return `["fail",{"message":"unknown command: ` + request + `"}]`
}

func (parser *wire_protocol_parser) step_matches(parameters json.RawMessage) string {
	var patterns map[string]string
	json.Unmarshal(parameters, &patterns)
	pattern := patterns["name_to_match"]
	id, matches, arguments := parser.steps.matching_step(pattern)
	if matches {
		arguments_string := build_arguments_string(arguments)
		return `["success",[{"id":"` + id + `","args":[` + arguments_string + `]}]]`
	}
	return `["success",[]]`
}

func (parser *wire_protocol_parser) snippet_text(parameters json.RawMessage) string {
	var snippet map[string]string
	json.Unmarshal(parameters, &snippet)
	name := strings.Replace(snippet["step_name"], "\"", "\\\"", -1)
	keyword := snippet["step_keyword"]
	snippet_text, _ := json.Marshal(`cucumber.` + keyword + `("` + name + `").Pending()`)
	return `["success",` + string(snippet_text) + `]`
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
		return `["fail",{"message":"` + message + `"}]`
	case pending:
		return `["pending"]`
	}
	return `["success"]`
}

func (parser *wire_protocol_parser) begin_scenario() string {
	parser.steps.begin_scenario()
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
