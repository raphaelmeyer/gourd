package main

import (
	"github.com/raphaelmeyer/gourd"
	"net"
)

type step_context struct {
	testee gourd.Cucumber
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		testee := gourd.NewCucumber(nil)
		return &step_context{testee}
	})

	cucumber.Given("no step implementation").Pass()

	cucumber.When("I run cucumber").Do(
		func(context interface{}) {
			stepContext, ok := context.(*step_context)
			if ok {
				stepContext.testee.Run()
				conn, err := net.Dial("tcp", "localhost:1847")
				if err != nil {
					panic("Wire server is not listening")
				}
				conn.Close()
			} else {
				panic("context")
			}
		})

	cucumber.When("a new scenario begins").Pending()

	cucumber.When("the scenario has a step").Pending()

	cucumber.When("the scenario ends").Pending()

	cucumber.Then("an undefined step is reported").Pending()

	cucumber.Run()
}
