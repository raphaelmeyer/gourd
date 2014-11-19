package main

import (
	"github.com/raphaelmeyer/gourd"
)

type gourd_context struct {
	testee gourd.Cucumber
	executed bool
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		scenario := &gourd_context{}
		scenario.testee = gourd.NewCucumber(func() interface{} {
			return nil
		})

		scenario.testee.Given("a step which is pending").Pending()
		scenario.testee.Given("a step which passes").Pass()
		scenario.testee.Given("a step which fails").Fail()
		scenario.testee.Given("a step with code").Do(
			func(context interface{}) {
				//scenario.executed = true
			})

		return scenario
	})

	cucumber.Given("a go wire server").Do(
		func(context interface{}) {
			scenario, _ := context.(*gourd_context)
			scenario.testee.SetPort(2345)
			go func() {
				scenario.testee.Run()
			}()
		})

	cucumber.Then("the code was executed").Do(
		func(context interface{}) {
			scenario, _ := context.(*gourd_context)
			if ! scenario.executed {
				panic("code was not executed")
			}
		})

	cucumber.Run()
}
