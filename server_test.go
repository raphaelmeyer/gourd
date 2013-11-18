package gourd

import (
	"net"
	"testing"
	"time"
	"bufio"
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
		t.Fatalf("Wire server is not listening: %s", err)
	}

	conn.Close()

	select {
	case <-done:
	case <-time.After(1000000): // 1s
		t.Error("Wire server did not exit.")
	}
}

type CommandParserMock struct {
	calls []string
}

func (parser * CommandParserMock) Parse(command string) {
	parser.calls = append(parser.calls, command)
}

func Test_the_wire_server_reads_a_one_line_command_and_sends_it_to_the_command_parser(t *testing.T) {
	parser := &CommandParserMock{}
	testee := WireServer{parser}
	done := make(chan bool)
	go func() {
		testee.Listen()
		done <- true
	}()

	conn, err := net.Dial("tcp", "localhost:1847")
	if err != nil {
		t.Fatalf("Wire server is not listening: %s", err)
	}

	writer := bufio.NewWriter(conn)

	command := "[\"begin_scenario\"]\n"
	_, err = writer.WriteString(command)
	if err != nil {
		t.Fatalf("Failed to send command to wire server: %s", err)
	}

	writer.Flush()
	conn.Close()

	select {
	case <-done:
	case <-time.After(1000000): // 1s
		t.Error("Wire server did not exit.")
	}

	if len(parser.calls) != 1 {
		t.Fatalf("Parse was called %d times, but expected once.", len(parser.calls))
	}

	if parser.calls[0] != command {
		t.Fatalf("Called to parse %s, but expected %s", parser.calls[0], command)
	}
}

