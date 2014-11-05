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

	cucumber.Given("A step that passes").Pass()

	cucumber.Given("A step that fails").Fail()

	cucumber.Run()
}
