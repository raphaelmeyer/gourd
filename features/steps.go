package main

import (
	"bufio"
	"fmt"
	"github.com/raphaelmeyer/gourd"
	"net"
	"time"
)

type gourd_context struct {
	testee   gourd.Cucumber
	conn     net.Conn
	done     chan bool
	contexts int
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		context := &gourd_context{}
		testee := gourd.NewCucumber(func() interface{} {
			context.contexts++
			return nil
		})
		context.testee = testee
		return context
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

	cucumber.Given("no given, when or then step").Pass()
	cucumber.When("cucumber runs the scenario").Pass()

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

	cucumber.When("a new scenario starts").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			writer := bufio.NewWriter(step_context.conn)
			writer.WriteString(`["begin_scenario"]` + "\n")
			writer.Flush()
			reader := bufio.NewReader(step_context.conn)
			response, _ := reader.ReadString('\n')
			if response != `["success"]`+"\n" {
				fmt.Println(response)
				panic("begin scenario failed")
			}
		})

	cucumber.When("the scenario ends").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			writer := bufio.NewWriter(step_context.conn)
			writer.WriteString(`["end_scenario"]` + "\n")
			writer.Flush()
			reader := bufio.NewReader(step_context.conn)
			response, _ := reader.ReadString('\n')
			if response != `["success"]`+"\n" {
				fmt.Println(response)
				panic("end scenario failed")
			}
		})

	cucumber.Then("a new context is created").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			if step_context.contexts != 1 {
				panic(fmt.Sprintf("Expected 1 created context, but actual value is %d", step_context.contexts))
			}
		})

	cucumber.Run()
}
