package gourd

type Step interface {
	Do(action func(context interface{}))
	Pass()
	Pending()
	Fail()
}

type gourd_step struct {
	pattern string
	action  func(context interface{})
}

func (step *gourd_step) Do(action func(context interface{})) {
	step.action = action
}

func (step *gourd_step) Pass() {
	step.action = func(context interface{}) {
	}
}

func (step *gourd_step) Pending() {
}

func (step *gourd_step) Fail() {
	step.action = func(context interface{}) {
		panic("")
	}
}
