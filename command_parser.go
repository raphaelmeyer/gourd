package gourd

import (
	"encoding/json"
)

type command_parser interface {
	parse() wire_command
}

type wire_command_parser struct {
}

func (parser *wire_command_parser) parse(command []byte) wire_command {
	var raw_command []json.RawMessage
	json.Unmarshal(command, &raw_command)

	var request string
	json.Unmarshal(raw_command[0], &request)

	switch request {
	case "begin_scenario":
		return &wire_command_begin_scenario{}
	case "end_scenario":
		return &wire_command_end_scenario{}
	case "step_matches":
		return &wire_command_step_matches{}
	}

	return nil
}
