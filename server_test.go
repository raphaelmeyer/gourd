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
	testee := &WireServer{}
	done := start_wireserver(testee)

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wire server is not listening.")

	conn.Close()

	assert_wireserver_exits(t, done)
}

func Test_wireserver_reads_and_parses_a_line(t *testing.T) {
	parser := &CommandParserMock{}
	testee := &WireServer{parser}
	done := start_wireserver(testee)

	command := "[\"begin_scenario\"]\n"
	parser.On("Parse", command).Return().Once()

	// Give the wire server some time to start accepting connection
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:1847")
	assert.Nil(t, err, "Wire server is not listening.")

	writer := bufio.NewWriter(conn)
	_, err = writer.WriteString(command)
	assert.Nil(t, err, "Failed to send command to wire server.")

	writer.Flush()
	conn.Close()

	assert_wireserver_exits(t, done)
	parser.Mock.AssertExpectations(t)
}

func start_wireserver(server *WireServer) chan bool {
	done := make(chan bool)
	go func() {
		server.Listen()
		done <- true
	}()
	return done
}

func assert_wireserver_exits(t *testing.T, done chan bool) {
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		assert.Fail(t, "Wire server did not exit.")
	}
}

type CommandParserMock struct {
	mock.Mock
}

func (parser *CommandParserMock) Parse(command string) {
	parser.Mock.Called(command)
}
