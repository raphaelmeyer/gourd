package gourd

type wire_command_parser struct {
}

func (parser *wire_command_parser) parse(command []byte) wire_command {
	return &wire_command_begin_scenario{}
}

