package main

import (
	"github.com/raphaelmeyer/gourd"
	"net"
	"time"
)

type gourd_context struct {
	testee gourd.Cucumber
	conn   net.Conn
	done   chan bool
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		testee := gourd.NewCucumber(nil)
		return &gourd_context{testee, nil, nil}
	})

	cucumber.Given("a wire server running on port 2345").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			step_context.testee.SetPort(2345)
			step_context.done = make(chan bool)
			go func() {
				step_context.testee.Run()
				step_context.done <- true
			}()
		})

	cucumber.When("cucumber connects to port 2345").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			step_context.conn, _ = net.Dial("tcp", "localhost:2345")
		})

	cucumber.When("cucumber closes the connection").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			step_context.conn.Close()
		})

	cucumber.Then("the wire server exits").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			select {
			case <-step_context.done:
			case <-time.After(time.Second):
				panic("Wire server still running")
			}
		})

	cucumber.Run()
}
