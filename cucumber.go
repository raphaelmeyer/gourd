package gourd

const DefaultPort string = ":1847"

type Cucumber interface {
	Given(pattern string) *Step
	When(pattern string) *Step
	Then(pattern string) *Step
	Assert(cond bool)
	Run()
}

func NewCucumber() Cucumber {
	steps := &cucumberSteps{}
	server := new_wire_server(steps)
	return &gourdCucumber{steps, server}
}

type gourdCucumber struct {
	steps  steps
	server wire_server
}

type Step struct {
}

func (cucumber *gourdCucumber) Given(pattern string) *Step {
	cucumber.steps.addStep(pattern)
	return &Step{}
}

func (cucumber *gourdCucumber) When(pattern string) *Step {
	return &Step{}
}

func (cucumber *gourdCucumber) Then(pattern string) *Step {
	return &Step{}
}

func (cucumber *gourdCucumber) Assert(cond bool) {
}

func (cucumber *gourdCucumber) Run() {
	cucumber.server.listen()
}

func (step *Step) Do(action func(context interface{})) {
}

func (step *Step) Pass() {
}

func (step *Step) Pending() {
}
