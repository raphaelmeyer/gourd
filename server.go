package gourd

import (
	"bufio"
	"net"
)

const DefaultPort string = ":1847"

type WireServer struct {
	parser Parser
}

func (server *WireServer) Listen() {
	listener, err := net.Listen("tcp", DefaultPort)
	if err != nil {
		panic(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)

	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		response := server.parser.Parse(command)
		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(response)
		writer.Flush()

	}

	listener.Close()
	conn.Close()
}
