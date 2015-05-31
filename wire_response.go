package gourd

type wire_response interface {
	encode() string
}

type generic_wire_response struct{}

func (response *generic_wire_response) encode() string {
	return ""
}

type command_parser interface {
	parse() wire_command
}
