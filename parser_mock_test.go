package gourd

import (
	"github.com/stretchr/testify/mock"
)

type parser_mock struct {
	mock.Mock
}

func (parser *parser_mock) parse(command []byte) string {
	args := parser.Mock.Called(command)
	return args.String(0)
}
