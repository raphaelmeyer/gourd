package gourd

type Cucumber struct {
}

type Step struct {
}

type Context interface {
}

func (cucumber *Cucumber) Given(step string) *Step {
	return new(Step)
}

func (cucumber *Cucumber) Expect(cond bool) {
}

func (cucumber *Cucumber) Start() {
	server := NewWireServer()
	server.Listen()
}

func (step *Step) Do(action func(constext *Context)) {
}
