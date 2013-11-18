package gourd

type Parser interface {
	Parse(command string)
}

type CommandParser struct {
}


