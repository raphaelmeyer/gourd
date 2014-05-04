package gourd

import (
	"fmt"
)

type step_result int

const (
	success step_result = iota
	pending
	fail
)

type steps interface {
	begin_scenario()
	matching_step(step string) (string, bool)
	add_step(pattern string) Step
	invoke_step(id string) (step_result, string)
}

type gourd_steps struct {
	new_context func() interface{}
	steps       map[string]*gourd_step
	id          int
}

func (steps *gourd_steps) begin_scenario() {
	if steps.new_context != nil {
		steps.new_context()
	}
}

func (steps *gourd_steps) matching_step(pattern string) (string, bool) {
	for id, step := range steps.steps {
		if step.pattern == pattern {
			return id, true
		}
	}
	return "", false
}

func (steps *gourd_steps) add_step(pattern string) Step {
	if steps.steps == nil {
		steps.steps = make(map[string]*gourd_step)
	}
	step := &gourd_step{pattern, nil}
	steps.steps[steps.nextId()] = step
	return step
}

func (steps *gourd_steps) invoke_step(id string) (result step_result, message string) {
	step, ok := steps.steps[id]
	if !ok {
		return fail, ""
	}

	if step.action == nil {
		return pending, ""
	}

	defer func() {
		if err := recover(); err != nil {
			result = fail
			message = err.(string)
		}
	}()

	step.action(nil)

	return success, ""
}

func (steps *gourd_steps) nextId() string {
	steps.id++
	return fmt.Sprintf("%d", steps.id)
}
