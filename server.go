package gourd

import (
	"bufio"
	"net"
)

const DefaultPort string = ":1847"

type WireServer struct {
}

func (*WireServer) Listen() {
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
		_, err = reader.ReadString('\n')
		if err != nil {
			break
		}
	}

	listener.Close()
	conn.Close()
}
