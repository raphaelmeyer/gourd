package gourd

import (
	"net"
	"testing"
	"time"
)

func Test_the_wire_server_listens_on_port_1847_and_exits_after_the_connection_is_closed(t *testing.T) {
	testee := new(WireServer)
	done := make(chan bool)
	go func() {
		testee.Listen()
		done <- true
	}()

	conn, err := net.Dial("tcp", "localhost:1847")
	if err != nil {
		t.Errorf("Wire server is not listening: %s", err)
	}

	conn.Close()

	select {
	case <-done:
	case <-time.After(1000000): // 1s
		t.Error("Wire server did not exit.")
	}
}
