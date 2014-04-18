package main

import (
	"github.com/raphaelmeyer/gourd"
	"net"
)

type stepContext struct {
	testee gourd.Cucumber
}

func main() {
	cucumber := gourd.NewCucumber()

	cucumber.Given("no step implementation").Pass()

	cucumber.When("I run cucumber").Do(
		func(context interface{}) {
			stepContext, ok := context.(stepContext)
			if ok {
				stepContext.testee.Run()
				conn, err := net.Dial("tcp", "localhost:1847")
				cucumber.Assert(err == nil)
				conn.Close()
			}
		})

	cucumber.When("a new scenario begins").Pending()

	cucumber.When("the scenario has a step").Pending()

	cucumber.When("the scenario ends").Pending()

	cucumber.Then("an undefined step is reported").Pending()

	cucumber.Run()
}
