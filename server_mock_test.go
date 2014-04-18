package gourd

import (
	"github.com/stretchr/testify/mock"
)

type server_mock struct {
	mock.Mock
}

func (server *server_mock) listen() {
	server.Mock.Called()
}
