package gourd

import "strings"

type Parser interface {
	Parse(command string) string
}

type CommandParser struct {
	steps IStepManager
}

func NewCommandParser() *CommandParser {
	steps := new(StepManager)
	return &CommandParser{steps}
}

func (parser *CommandParser) Parse(command string) string {
	if strings.Contains(command, `"step_matches"`) {
		matches, _ := parser.steps.MatchingStep("defined step")
		if matches {
			return `["success",[{"id":"1", "args":[]}]]` + "\n"
		}
		return `["success",[]]` + "\n"
	}
	return `["success"]` + "\n"
}

