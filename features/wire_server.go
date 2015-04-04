package main

import (
	"github.com/raphaelmeyer/gourd"
)

type gourd_context struct {
	testee         gourd.Cucumber
	executed       bool
	matched_number int
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
			func(context interface{}, arguments gourd.Arguments) {
				scenario.executed = true
			})

		return scenario
	})

	cucumber.Given("a step with pattern \"failure step\" that fails with message \"failure message\"").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*gourd_context)
			scenario.testee.Given("failure step").Do(
				func(context interface{}, arguments gourd.Arguments) {
					panic("failure message")
				})
		})

	cucumber.Given("a go wire server").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*gourd_context)
			scenario.testee.SetPort(2345)
			go func() {
				scenario.testee.Run()
			}()
		})

	cucumber.Then("the code was executed").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*gourd_context)
			if !scenario.executed {
				panic("code was not executed")
			}
		})

	cucumber.Given("step with pattern \"\\^a number \\(\\\\d\\+\\)\\$\"").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*gourd_context)
			scenario.testee.Given("^a number (\\d+)$").Do(
				func(context interface{}, arguments gourd.Arguments) {
					scenario.matched_number = arguments.Int(0)
				})
		})

	cucumber.Then("number 1234 is passed to the matching step").Do(
		func(context interface{}, arguments gourd.Arguments) {
			scenario, _ := context.(*gourd_context)
			if scenario.matched_number != 1234 {
				panic("did not match expected number")
			}
		})

	cucumber.Run()
}
