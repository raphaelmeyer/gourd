package gourd

type wire_response interface {
	encode() string
}

type wire_command interface {
	execute() wire_response
}

type generic_wire_command struct{}
type generic_wire_response struct{}

type wire_command_begin_scenario struct {
}

func (command *wire_command_begin_scenario) execute() wire_response {
	return &generic_wire_response{}
}

func parse_wire_command(command []byte) wire_command {
	return &wire_command_begin_scenario{}
}




