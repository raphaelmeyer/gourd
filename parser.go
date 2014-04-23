package gourd

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type parser interface {
	parse(command []byte) string
}

type wire_protocol_parser struct {
	steps steps
}

func (parser *wire_protocol_parser) parse(command []byte) string {
	var data []interface{}
	_ = json.Unmarshal(command, &data)
	request := data[0].(string)
	switch request {
	case "step_matches":
		return parser.step_matches(data[1])
	case "begin_scenario":
		return `["success"]` + "\n"
	case "end_scenario":
		return `["success"]` + "\n"
	case "snippet_text":
		return parser.snippet_text(data[1])
	case "invoke":
		return parser.invoke(data[1])
	}
	return `["fail",{"message":"unknown command: ` + request + `"}]` + "\n"
}

func (parser *wire_protocol_parser) step_matches(parameters interface{}) string {
	pattern := parameters.(map[string]interface{})["name_to_match"].(string)
	matches, id := parser.steps.matching_step(pattern)
	if matches {
		return fmt.Sprintf(`["success",[{"id":"%d", "args":[]}]]`+"\n", id)
	}
	return `["success",[]]` + "\n"
}

func (parser *wire_protocol_parser) snippet_text(parameters interface{}) string {
	snippet := parameters.(map[string]interface{})
	name := snippet["step_name"].(string)
	keyword := snippet["step_keyword"].(string)
	return `["success","cucumber.` + keyword + `(\"` + name + `\").Pending()\n"]` + "\n"
}

func (parser *wire_protocol_parser) invoke(parameters interface{}) string {
	invoke := parameters.(map[string]interface{})
	id_string := invoke["id"].(string)
	id, _ := strconv.Atoi(id_string)
	parser.steps.invoke_step(id)
	return `["success"]` + "\n"
}
