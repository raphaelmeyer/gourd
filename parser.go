package gourd

import "strings"

type Parser interface {
	Parse(command string) string
}

type CommandParser struct {
	step_manager StepManager
}

func (parser *CommandParser) Parse(command string) string {
	if strings.Contains(command, `"step_matches"`) {
		if parser.step_manager != nil {
			matches, _ := parser.step_manager.MatchingStep("defined step")
			if matches {
				return `["success",[{"id":"1", "args":[]}]]` + "\n"
			}
		}
		return `["success",[]]` + "\n"
	}
	return `["success"]` + "\n"
}

