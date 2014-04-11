package gourd

type Cucumber struct {
}

type Step struct {
}

type Context interface {
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

func (cucumber *Cucumber) Expect(cond bool) {
}

func (cucumber *Cucumber) Start() {
	server := newWireServer()
	server.Listen()
}

func (step *Step) Do(action func(constext *Context)) {
}

func (step *Step) Pending() {
}
