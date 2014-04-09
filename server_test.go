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
	testee := &wireServer{}
	done := startWireServer(testee)

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wire server is not listening.")

	conn.Close()

	assertWireServerExits(t, done)
}

func Test_wireserver_reads_and_parses_a_line(t *testing.T) {
	parser := &commandParserMock{}
	testee := &wireServer{parser}
	done := startWireServer(testee)

	command := []byte("[\"begin_scenario\"]\n")
	parser.On("parse", command).Return("").Once()

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wireserver is not listening.")

	writer := bufio.NewWriter(conn)
	_, err = writer.Write(command)
	assert.Nil(t, err, "Failed to send command to wire server.")

	writer.Flush()
	conn.Close()

	assertWireServerExits(t, done)
	parser.Mock.AssertExpectations(t)
}

func Test_wireserver_reads_and_parses_next_line_after_processing_first_one(t *testing.T) {
	parser := &commandParserMock{}
	testee := &wireServer{parser}
	done := startWireServer(testee)

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wireserver is not listening.")

	// First command
	command := []byte("[\"begin_scenario\"]\n")
	parser.On("parse", command).Return("").Once()
	writer := bufio.NewWriter(conn)
	_, err = writer.Write(command)
	assert.Nil(t, err, "Failed to send command to wire server.")

	writer.Flush()

	// Next command
	command = []byte("[\"end_scenario\"]\n")
	parser.On("parse", command).Return("").Once()
	_, err = writer.Write(command)
	assert.Nil(t, err, "Failed to send command to wire server.")

	writer.Flush()
	conn.Close()

	assertWireServerExits(t, done)
	parser.Mock.AssertExpectations(t)
}

func Test_wireserver_writes_response_from_parser(t *testing.T) {
	parser := &commandParserMock{}
	testee := &wireServer{parser}
	done := startWireServer(testee)

	command := []byte("[\"begin_scenario\"]\n")
	response := "[\"success\"]\n"
	parser.On("parse", command).Return(response).Once()

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wireserver is not listening.")

	writer := bufio.NewWriter(conn)
	_, err = writer.Write(command)
	assert.Nil(t, err, "Failed to send command to wire server.")
	writer.Flush()

	assert_wireserver_responds(t, conn, response)

	conn.Close()

	assertWireServerExits(t, done)
	parser.Mock.AssertExpectations(t)
}

func startWireServer(server *wireServer) chan bool {
	done := make(chan bool)
	go func() {
		server.Listen()
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

type commandParserMock struct {
	mock.Mock
}

func (parser *commandParserMock) parse(command []byte) string {
	args := parser.Mock.Called(command)
	return args.String(0)
}
