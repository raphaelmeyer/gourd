package gourd

type StepManager interface {
	MatchingStep(pattern string) (bool, int)
}

