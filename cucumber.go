package gourd

type Cucumber struct {
}

type Step struct {
}

func (cucumber *Cucumber) Given(step string) *Step {
	return &Step{}
}

func (cucumber *Cucumber) When(step string) *Step {
	return &Step{}
}

func (cucumber *Cucumber) Then(step string) *Step {
	return &Step{}
}

func (cucumber *Cucumber) Assert(cond bool) {
}

func (cucumber *Cucumber) Start() {
	server := newWireServer()
	server.Listen()
}

func (step *Step) Do(action func(context interface{})) {
}

func (step *Step) Pass() {
}

func (step *Step) Pending() {
}
