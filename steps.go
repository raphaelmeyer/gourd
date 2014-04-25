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
	matching_step(step string) (bool, string)
	add_step(pattern string) Step
	invoke_step(id string) step_result
}

type gourd_steps struct {
	steps map[string]*gourd_step
	id    int
}

func (steps *gourd_steps) matching_step(pattern string) (bool, string) {
	for id, step := range steps.steps {
		if step.pattern == pattern {
			return true, id
		}
	}
	return false, ""
}

func (steps *gourd_steps) add_step(pattern string) Step {
	if steps.steps == nil {
		steps.steps = make(map[string]*gourd_step)
	}
	step := &gourd_step{pattern, nil}
	steps.steps[steps.nextId()] = step
	return step
}

func (steps *gourd_steps) invoke_step(id string) (result step_result) {
	step, ok := steps.steps[id]
	if !ok {
		return fail
	}

	if step.action == nil {
		return pending
	}

	defer func() {
		if recover() != nil {
			result = fail
		}
	}()

	step.action(nil)

	return success
}

func (steps *gourd_steps) nextId() string {
	steps.id++
	return fmt.Sprintf("%d", steps.id)
}
