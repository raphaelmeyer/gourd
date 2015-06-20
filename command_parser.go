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
		return parse_step_matches(raw_command[1])
	}

	return nil
}

func parse_step_matches(raw_arguments json.RawMessage) wire_command {
	type name_to_match struct {
		Pattern string `json:"name_to_match"`
	}

	var arguments name_to_match
	json.Unmarshal(raw_arguments, &arguments)

	return &wire_command_step_matches{arguments.Pattern}
}
