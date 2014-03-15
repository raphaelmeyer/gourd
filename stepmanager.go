package gourd

type IStepManager interface {
	MatchingStep(step string) (bool, int)
	AddStep(pattern string)
}

type StepManager struct {
}

func (steps *StepManager) MatchingStep(step string) (bool, int) {
	if step == "pattern" {
		return true, 1
	}
	return false, 0
}

func (steps *StepManager) AddStep(pattern string) {

}
