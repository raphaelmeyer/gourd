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
	matching_step(step string) (string, bool, []capturing_group)
	add_step(pattern string) Step
	invoke_step(id string, arguments []string) (step_result, string)
}

type capturing_group struct {
	value    string
	position uint
}

type gourd_steps struct {
	new_context func() interface{}
	context     interface{}
	steps       map[string]*gourd_step
	id          int
}

func (steps *gourd_steps) begin_scenario() {
	if steps.new_context != nil {
		steps.context = steps.new_context()
	}
}

func (steps *gourd_steps) matching_step(pattern string) (string, bool, []capturing_group) {
	for id, step := range steps.steps {
		if step.regex.MatchString(pattern) {
			arguments := []capturing_group{}
			submatches := step.regex.FindStringSubmatch(pattern)
			positions := step.regex.FindStringSubmatchIndex(pattern)
			for i, submatch := range submatches {
				if i > 0 {
					position := uint(positions[2*i])
					arguments = append(arguments, capturing_group{value: submatch, position: position})
				}
			}
			return id, true, arguments
		}
	}
	return "", false, []capturing_group{}
}

func (steps *gourd_steps) add_step(pattern string) Step {
	if steps.steps == nil {
		steps.steps = make(map[string]*gourd_step)
	}
	step := new_step(pattern)
	steps.steps[steps.nextId()] = step
	return step
}

func (steps *gourd_steps) invoke_step(id string, arguments []string) (result step_result, message string) {
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
			message = fmt.Sprintf("%v", err)
		}
	}()

	step.action(steps.context, &gourd_arguments{arguments})

	return success, ""
}

func (steps *gourd_steps) nextId() string {
	steps.id++
	return fmt.Sprintf("%d", steps.id)
}
