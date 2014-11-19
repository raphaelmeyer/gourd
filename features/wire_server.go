package main

import (
	"github.com/raphaelmeyer/gourd"
)

type gourd_context struct {
	testee gourd.Cucumber
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		context := &gourd_context{}
		context.testee = gourd.NewCucumber(func() interface{} {
			return nil
		})

		context.testee.Given("a step which is pending").Pending()
		context.testee.Given("a step which passes").Pass()
		context.testee.Given("a step which fails").Fail()

		return context
	})

	cucumber.Given("a go wire server").Do(
		func(context interface{}) {
			scenario, _ := context.(*gourd_context)
			scenario.testee.SetPort(2345)
			go func() {
				scenario.testee.Run()
			}()
		})


	cucumber.Run()
}
