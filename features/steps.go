package main

import (
  "github.com/raphaelmeyer/gourd"
)

func main() {
  cucumber := &gourd.Cucumber{}

  cucumber.Given(`a scenario with step "arbitrary step"`)

  cucumber.When(`I run cucumber`)

  cucumber.Then(`it should return step "arbitrary step" is undefined`)

  cucumber.Start();
}

