package gourd

type Step interface {
	Do(action func(context interface{}))
	Pass()
	Pending()
}

type gourd_step struct {
	pattern string
}

func (step *gourd_step) Do(action func(context interface{})) {
}

func (step *gourd_step) Pass() {
}

func (step *gourd_step) Pending() {
}
