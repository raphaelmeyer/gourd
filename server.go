package gourd

import (
	"bufio"
	"fmt"
	"net"
)

type wire_server interface {
	listen(port uint)
}

type gourd_wire_server struct {
	parser parser
}

func new_wire_server(steps steps) wire_server {
	parser := &wire_protocol_parser{steps}
	return &gourd_wire_server{parser}
}

func (server *gourd_wire_server) listen(port uint) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port", port)

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Bytes()

		response := server.parser.parse(command)
		writer := bufio.NewWriter(conn)
		writer.WriteString(response + "\n")
		writer.Flush()
	}
}
