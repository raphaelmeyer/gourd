package gourd

type steps interface {
	matching_step(step string) (bool, int)
	add_step(pattern string) Step
}

type gourd_steps struct {
	steps map[int]*gourd_step
	id    int
}

func (steps *gourd_steps) matching_step(pattern string) (bool, int) {
	for id, step := range steps.steps {
		if step.pattern == pattern {
			return true, id
		}
	}
	return false, 0
}

func (steps *gourd_steps) add_step(pattern string) Step {
	if steps.steps == nil {
		steps.steps = make(map[int]*gourd_step)
	}
	step := &gourd_step{pattern}
	steps.steps[steps.nextId()] = step
	return step
}

func (steps *gourd_steps) nextId() int {
	steps.id++
	return steps.id
}
