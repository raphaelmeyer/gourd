package gourd

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net"
	"testing"
	"time"
)

func Test_wireserver_accepts_one_connection_on_port_1847(t *testing.T) {
	testee := &gourd_wire_server{}
	done := startWireServer(testee)

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wire server is not listening.")

	conn.Close()

	assertWireServerExits(t, done)
}

func Test_wireserver_forwards_commands_to_parser_without_newlines(t *testing.T) {
	parser := &parser_mock{}
	testee := &gourd_wire_server{parser}
	done := startWireServer(testee)

	expected_command := []byte(`["begin_scenario"]`)
	parser.On("parse", expected_command).Return("").Once()

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wireserver is not listening.")

	command := append(expected_command, '\n')

	writer := bufio.NewWriter(conn)
	_, err = writer.Write(command)
	assert.Nil(t, err, "Failed to send command to wire server.")

	writer.Flush()
	conn.Close()

	assertWireServerExits(t, done)
	parser.Mock.AssertExpectations(t)
}

func Test_wireserver_reads_and_parses_line_by_line(t *testing.T) {
	parser := &parser_mock{}
	testee := &gourd_wire_server{parser}
	done := startWireServer(testee)

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wireserver is not listening.")

	// First command
	command := []byte(`["begin_scenario"]`)
	parser.On("parse", command).Return("").Once()
	writer := bufio.NewWriter(conn)
	_, err = writer.Write(append(command, '\n'))
	assert.Nil(t, err, "Failed to send command to wire server.")

	writer.Flush()

	// Next command
	command = []byte(`["end_scenario"]`)
	parser.On("parse", command).Return("").Once()
	_, err = writer.Write(append(command, '\n'))
	assert.Nil(t, err, "Failed to send command to wire server.")

	writer.Flush()
	conn.Close()

	assertWireServerExits(t, done)
	parser.Mock.AssertExpectations(t)
}

func Test_wireserver_sends_response_from_parser_including_newline(t *testing.T) {
	parser := &parser_mock{}
	testee := &gourd_wire_server{parser}
	done := startWireServer(testee)

	response := `["success"]`
	parser.On("parse", mock.Anything).Return(response).Once()

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wireserver is not listening.")

	command := []byte(`["begin_scenario"]`)
	writer := bufio.NewWriter(conn)
	_, err = writer.Write(append(command, '\n'))
	assert.Nil(t, err, "Failed to send command to wire server.")
	writer.Flush()

	expected_response := response + "\n"
	assert_wireserver_responds(t, conn, expected_response)

	conn.Close()

	assertWireServerExits(t, done)
	parser.Mock.AssertExpectations(t)
}

func startWireServer(server *gourd_wire_server) chan bool {
	done := make(chan bool)
	go func() {
		server.listen()
		done <- true
	}()
	return done
}

func assertWireServerExits(t *testing.T, done chan bool) {
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		assert.Fail(t, "Wireserver did not exit.")
	}
}

func assert_wireserver_responds(t *testing.T, conn net.Conn, response string) {
	done := make(chan bool)
	go func() {
		reader := bufio.NewReader(conn)
		line, err := reader.ReadString('\n')
		assert.Nil(t, err, "Failed to read response from wire server.")
		assert.Equal(t, response, line)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		assert.Fail(t, "Wireserver did not respond.")
	}
}
