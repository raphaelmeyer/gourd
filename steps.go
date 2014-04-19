package gourd

type steps interface {
	matchingStep(step string) (bool, int)
	add_step(pattern string) Step
}

type cucumberSteps struct {
	steps map[int]*gourd_step
	id    int
}

func (steps *cucumberSteps) matchingStep(pattern string) (bool, int) {
	for id, step := range steps.steps {
		if step.pattern == pattern {
			return true, id
		}
	}
	return false, 0
}

func (steps *cucumberSteps) add_step(pattern string) Step {
	if steps.steps == nil {
		steps.steps = make(map[int]*gourd_step)
	}
	step := &gourd_step{pattern}
	steps.steps[steps.nextId()] = step
	return step
}

func (steps *cucumberSteps) nextId() int {
	steps.id++
	return steps.id
}
