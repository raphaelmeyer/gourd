package gourd

type Arguments interface {
	String(index uint) string
	Int(index uint) int
	Uint(index uint) uint
}

type gourd_arguments struct {
}

func (arguments *gourd_arguments) String(index uint) string {
	return ""
}

func (arguments *gourd_arguments) Int(index uint) int {
	return 0
}

func (arguments *gourd_arguments) Uint(index uint) uint {
	return 0
}

type argument struct {
	value    string
	position uint
}
