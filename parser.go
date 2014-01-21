package gourd

import (
	"encoding/json"
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
	if (request == "step_matches") {
		pattern := data[1].(map[string]interface{})["name_to_match"].(string)
		matches, _ := parser.steps.MatchingStep(pattern)
		if matches {
			return `["success",[{"id":"1", "args":[]}]]` + "\n"
		}
		return `["success",[]]` + "\n"
	}
	return `["success"]` + "\n"
}

