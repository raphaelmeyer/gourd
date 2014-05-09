package gourd

const DefaultPort uint = 1847

type Cucumber interface {
	Given(pattern string) Step
	When(pattern string) Step
	Then(pattern string) Step
	Run()
	SetPort(port uint)
}

func NewCucumber(new_context func() interface{}) Cucumber {
	steps := &gourd_steps{}
	steps.new_context = new_context
	server := new_wire_server(steps)
	return &gourd_cucumber{steps, server, DefaultPort}
}

type gourd_cucumber struct {
	steps  steps
	server wire_server
	port   uint
}

func (cucumber *gourd_cucumber) Given(pattern string) Step {
	return cucumber.steps.add_step(pattern)
}

func (cucumber *gourd_cucumber) When(pattern string) Step {
	return cucumber.steps.add_step(pattern)
}

func (cucumber *gourd_cucumber) Then(pattern string) Step {
	return cucumber.steps.add_step(pattern)
}

func (cucumber *gourd_cucumber) Run() {
	cucumber.server.listen(cucumber.port)
}

func (cucumber *gourd_cucumber) SetPort(port uint) {
	cucumber.port = port
}
