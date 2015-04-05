/*
Gourd is a wire protocol implementation for Cucumber.
Gourd allows to write step definitions in go.

You will need to install ruby and Cucumber to use Gourd.
Installing Cucumber on top of ruby is usually as simple as "gem install cucumber".

Visit http://cukes.info/ to learn more about Cucumber.

Usage

In the directory from which you wish to run cucumber create sub-directories "features/step_definitions/".
Add the wire server configuration to "features/step_definitions/gourd.wire":

	host: localhost
	port: 1847

Put your .feature-files into "features/".

The main of the wire server can be put in any location.
The following example will assume "features/step_definitions/steps.go".
This may work well for smaller projects.
But for easier navigation and readablity,
its own folder with step definitions split across multiple files may be more appropriate.

A minimal main creates a wire server and runs it:

	package main

	import (
		"github.com/raphaelmeyer/gourd"
	)

	func main() {
		cucumber := gourd.NewCucumber(func() interface{} {
			return nil
		})
		cucumber.Run()
	}

Run the wire server...
	$ go run features/step_definitions/steps.go

...then start cucumber:
	$ cucumber

You need at least one feature scenario.
If there is no scenario defined then cucumber will not connect to the wire server and
the server keeps waiting for a connection.

The function passed to NewCucumber is used to create a new context for each scenario.
It will be called before running a scenario and its return value is passed to all steps.

Example

The feature scenario:

	Feature: Evaluate an input

		Scenario: A single input

			Given I enter 7
			And I press evaluate
			Then the result should be 42

And the wire server with the step definitions:

	type my_context struct {
		testee Testee
	}

	func main() {
		cucumber := gourd.NewCucumber(func() interface{} {
			return &my_context{}
		})

		cucumber.Given("^I enter (\\d+)$").Do(
			func(context interface{}, arguments gourd.Arguments) {
				scenario, _ := context.(*my_context)
				scenario.testee.Enter(arguments.Int(0))
			})

		cucumber.When("^I press evaluate$").Do(
			func(context interface{}, arguments gourd.Arguments) {
				scenario, _ := context.(*my_context)
				scenario.testee.Evaluate()
			})

		cucumber.Then("^the result should be (\\d+)$").Do(
			func(context interface{}, arguments gourd.Arguments) {
				scenario, _ := context.(*my_context)
				if scenario.testee.Result() != arguments.Int(0) {
					panic("Wrong result")
				}
			})

		cucumber.Run()
	}
*/
package gourd
