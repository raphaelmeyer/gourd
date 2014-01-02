package gourd

type Parser interface {
	Parse(command string) string
}

type CommandParser struct {
}

func (parser *CommandParser) Parse(command string) string {
	return ""
}

