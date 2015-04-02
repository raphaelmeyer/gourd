package gourd

import (
	"regexp"
)

type Step interface {
	Do(action func(context interface{}, args ...interface{}))
	Pass()
	Pending()
	Fail()
}

type gourd_step struct {
	regex  *regexp.Regexp
	action func(context interface{}, args ...interface{})
}

func (step *gourd_step) Do(action func(context interface{}, args ...interface{})) {
	step.action = action
}

func (step *gourd_step) Pass() {
	step.action = func(context interface{}, args ...interface{}) {
	}
}

func (step *gourd_step) Pending() {
}

func (step *gourd_step) Fail() {
	step.action = func(context interface{}, args ...interface{}) {
		panic("")
	}
}

func new_step(pattern string) *gourd_step {
	regex := regexp.MustCompile(pattern)
	return &gourd_step{regex, nil}
}
