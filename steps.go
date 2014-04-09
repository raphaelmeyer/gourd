package gourd

type steps interface {
	matchingStep(step string) (bool, int)
	addStep(pattern string)
}

type cucumberSteps struct {
	steps map[int]string
	id    int
}

func (steps *cucumberSteps) matchingStep(step string) (bool, int) {
	for id, pattern := range steps.steps {
		if step == pattern {
			return true, id
		}
	}
	return false, 0
}

func (steps *cucumberSteps) addStep(pattern string) {
	if steps.steps == nil {
		steps.steps = make(map[int]string)
	}
	steps.steps[steps.nextId()] = pattern
}

func (steps *cucumberSteps) nextId() int {
	steps.id++
	return steps.id
}
