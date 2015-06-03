package gourd

type wire_command interface {
	execute(steps steps) wire_response
}

type wire_command_begin_scenario struct {
}

func (command *wire_command_begin_scenario) execute(steps steps) wire_response {
	steps.begin_scenario()
	return &generic_wire_response{}
}

type wire_command_end_scenario struct {
}

func (command *wire_command_end_scenario) execute(steps steps) wire_response {
	return &generic_wire_response{}
}
