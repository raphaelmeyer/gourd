package gourd

import (
	"strconv"
)

type Arguments interface {
	String(index uint) string
	Int(index uint) int
}

type gourd_arguments struct {
	values []string
}

func (arguments *gourd_arguments) String(index uint) string {
	return arguments.values[index]
}

func (arguments *gourd_arguments) Int(index uint) int {
	value, err := strconv.ParseInt(arguments.values[index], 10, 0)
	if err != nil {
		panic(err.Error())
	}
	return int(value)
}
