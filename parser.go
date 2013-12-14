package gourd

type Parser interface {
	Parse(command string) string
}

type CommandParser struct {
}


