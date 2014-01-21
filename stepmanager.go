package gourd

type IStepManager interface {
	MatchingStep(pattern string) (bool, int)
}

type StepManager struct {
}

func (steps * StepManager) MatchingStep(pattern string) (bool, int) {
	return false, 0
}

