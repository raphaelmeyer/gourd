package main

import (
	"github.com/raphaelmeyer/gourd"
	"net"
)

type gourd_context struct {
	testee gourd.Cucumber
	conn   net.Conn
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		testee := gourd.NewCucumber(nil)
		testee.SetPort(2847)

		return &gourd_context{testee, nil}
	})

	cucumber.Given("a wire server running on port 1847").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			go step_context.testee.Run()
		})

	cucumber.Then("cucumber can connect to port 1847").Do(
		func(context interface{}) {
			step_context, _ := context.(*gourd_context)
			var err error
			step_context.conn, err = net.Dial("tcp", "localhost:2847")
			if err != nil {
				panic("Wire server is not listening")
			}
		})

	cucumber.Given("no step implementation").Pass()

	cucumber.When("I run cucumber").Pending()

	cucumber.When("a new scenario begins").Pending()

	cucumber.When("the scenario has a step").Pending()

	cucumber.When("the scenario ends").Pending()

	cucumber.Then("an undefined step is reported").Pending()

	cucumber.Run()
}
