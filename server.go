package gourd

import (
	"bufio"
	"fmt"
	"net"
)

type wire_server interface {
	listen()
}

type gourd_wire_server struct {
	parser parser
}

func new_wire_server(steps steps) wire_server {
	parser := &commandParser{steps}
	return &gourd_wire_server{parser}
}

func (server *gourd_wire_server) listen() {
	listener, err := net.Listen("tcp", DefaultPort)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port", DefaultPort)

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		command, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		response := server.parser.parse(command)
		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(response)
		writer.Flush()
	}
}
