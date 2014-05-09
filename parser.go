package gourd

import (
	"encoding/json"
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
	id, matches := parser.steps.matching_step(pattern)
	if matches {
		return `["success",[{"id":"` + id + `","args":[]}]]`
	}
	return `["success",[]]`
}

func (parser *wire_protocol_parser) snippet_text(parameters interface{}) string {
	snippet := parameters.(map[string]interface{})
	name := snippet["step_name"].(string)
	keyword := snippet["step_keyword"].(string)
	return `["success","cucumber.` + keyword + `(\"` + name + `\").Pending()"]`
}

func (parser *wire_protocol_parser) invoke(parameters interface{}) string {
	invoke := parameters.(map[string]interface{})
	id := invoke["id"].(string)
	result, message := parser.steps.invoke_step(id)
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
