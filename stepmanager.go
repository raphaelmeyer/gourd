package gourd

type IStepManager interface {
	MatchingStep(step string) (bool, int)
	AddStep(pattern string)
}

type StepManager struct {
	steps map[int]string
	id    int
}

func (steps *StepManager) MatchingStep(step string) (bool, int) {
	for id, pattern := range steps.steps {
		if step == pattern {
			return true, id
		}
	}
	return false, 0
}

func (steps *StepManager) AddStep(pattern string) {
	if steps.steps == nil {
		steps.steps = make(map[int]string)
	}
	steps.steps[steps.nextId()] = pattern
}

func (steps *StepManager) nextId() int {
	steps.id++
	return steps.id
}
