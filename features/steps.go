package main

import (
	"github.com/raphaelmeyer/gourd"
	"net"
)

type gourd_context struct {
	testee gourd.Cucumber
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		testee := gourd.NewCucumber(nil)
		testee.SetPort(2847)

		return &gourd_context{testee}
	})

	cucumber.Given("no step implementation").Pass()

	cucumber.When("I run cucumber").Do(
		func(context interface{}) {
			step_context, ok := context.(*gourd_context)
			if ok {
				go step_context.testee.Run()
				conn, err := net.Dial("tcp", "localhost:2847")
				if err != nil {
					panic("Wire server is not listening")
				}
				conn.Close()
			} else {
				panic("Context missing")
			}
		})

	cucumber.When("a new scenario begins").Pending()

	cucumber.When("the scenario has a step").Pending()

	cucumber.When("the scenario ends").Pending()

	cucumber.Then("an undefined step is reported").Pending()

	cucumber.Run()
}
