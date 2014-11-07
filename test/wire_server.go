package main

import (
	"github.com/raphaelmeyer/gourd"
)

type gourd_context struct {
}

func main() {
	cucumber := gourd.NewCucumber(func() interface{} {
		context := &gourd_context{}
		return context
	})

//	cucumber.Given("A step with context").Do(
//		func(context interface{}) {
//			step_context, _ := context.(*gourd_context)
//		})

	cucumber.Given("a step which is pending").Pending()

	cucumber.Given("a step which passes").Pass()

	cucumber.Given("a step which fails").Fail()

	cucumber.Run()
}
