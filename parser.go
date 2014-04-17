package gourd

import (
	"encoding/json"
	"fmt"
)

type parser interface {
	parse(command []byte) string
}

type commandParser struct {
	steps steps
}

func newCommandParser() *commandParser {
	steps := &cucumberSteps{}
	return &commandParser{steps}
}

func (parser *commandParser) parse(command []byte) string {
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
	}
	return `["fail",{"message":"unknown command"}]` + "\n"
}

func (parser *commandParser) step_matches(parameters interface{}) string {
	pattern := parameters.(map[string]interface{})["name_to_match"].(string)
	matches, id := parser.steps.matchingStep(pattern)
	if matches {
		return fmt.Sprintf(`["success",[{"id":"%d", "args":[]}]]`+"\n", id)
	}
	return `["success",[]]` + "\n"
}

func (parser *commandParser) snippet_text(parameters interface{}) string {
	snippet := parameters.(map[string]interface{})
	name := snippet["step_name"].(string)
	keyword := snippet["step_keyword"].(string)
	return `["success","cucumber.` + keyword + `(\"` + name + `\").Pending()\n"]` + "\n"
}
