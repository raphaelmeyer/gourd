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
		return `["success",[]]` + "\n"
	}
	return `["success"]` + "\n"
}
