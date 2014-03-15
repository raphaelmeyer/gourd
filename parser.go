package gourd

import (
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse(command []byte) string
}

type CommandParser struct {
	steps IStepManager
}

func NewCommandParser() *CommandParser {
	steps := new(StepManager)
	return &CommandParser{steps}
}

func (parser *CommandParser) Parse(command []byte) string {
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
		return `["success","cucumber.Given(\"Step\").Pending()\n"]` + "\n"
	}
	return `["fail",{"message":"unknown command"}]` + "\n"
}

func (parser *CommandParser) step_matches(parameters interface{}) string {
	pattern := parameters.(map[string]interface{})["name_to_match"].(string)
	matches, id := parser.steps.MatchingStep(pattern)
	if matches {
		return fmt.Sprintf(`["success",[{"id":"%d", "args":[]}]]`+"\n", id)
	}
	return `["success",[]]` + "\n"
}
