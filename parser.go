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
	var data []interface{}
	err := json.Unmarshal(command, &data)
	if err != nil {
		return `["fail",{"message":"invalid command"}]`
	}

	return parser.evaluate(data)
}

func (parser *wire_protocol_parser) evaluate(command []interface{}) string {
	request := command[0].(string)
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

func (parser *wire_protocol_parser) step_matches(parameters interface{}) string {
	pattern := parameters.(map[string]interface{})["name_to_match"].(string)
	id, matches, arguments := parser.steps.matching_step(pattern)
	if matches {
		arguments_string := build_arguments_string(arguments)
		return `["success",[{"id":"` + id + `","args":[` + arguments_string + `]}]]`
	}
	return `["success",[]]`
}

func (parser *wire_protocol_parser) snippet_text(parameters interface{}) string {
	snippet := parameters.(map[string]interface{})
	name := strings.Replace(snippet["step_name"].(string), "\"", "\\\"", -1)
	keyword := snippet["step_keyword"].(string)
	snippet_text, _ := json.Marshal(`cucumber.` + keyword + `("` + name + `").Pending()`)
	return `["success",` + string(snippet_text) + `]`
}

func (parser *wire_protocol_parser) invoke(parameters interface{}) string {
	invoke := parameters.(map[string]interface{})
	id := invoke["id"].(string)
	args := invoke["args"].([]interface{})
	arguments := make([]string, len(args))
	for i, argument := range args {
		arguments[i] = argument.(string)
	}
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
