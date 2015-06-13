package gourd

type wire_command interface {
	execute(steps steps) wire_response
}

type wire_command_begin_scenario struct {
}

func (command *wire_command_begin_scenario) execute(steps steps) wire_response {
	steps.begin_scenario()
	return &wire_response_success{}
}

type wire_command_end_scenario struct {
}

func (command *wire_command_end_scenario) execute(steps steps) wire_response {
	return &wire_response_success{}
}

type wire_command_step_matches struct {
}

func (command *wire_command_step_matches) execute(steps steps) wire_response {
	return nil
}
