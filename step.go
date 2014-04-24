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

func new_step(pattern string) *gourd_step {
	return &gourd_step{pattern, nil}
}

func (step *gourd_step) Do(action func(context interface{})) {
}

func (step *gourd_step) Pass() {
	step.action = func(context interface{}) {
	}
}

func (step *gourd_step) Pending() {
}

func (step *gourd_step) Fail() {
}

func (step *gourd_step) is_pending() bool {
	return (step.action == nil)
}
