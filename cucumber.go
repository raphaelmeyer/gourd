package gourd

const DefaultPort string = ":1847"

type Cucumber interface {
	Given(pattern string) Step
	When(pattern string) Step
	Then(pattern string) Step
	Run()
}

func NewCucumber() Cucumber {
	steps := &gourd_steps{}
	server := new_wire_server(steps)
	return &gourd_cucumber{steps, server}
}

type gourd_cucumber struct {
	steps  steps
	server wire_server
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
	cucumber.server.listen()
}
