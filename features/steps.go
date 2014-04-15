package main

import (
  "github.com/raphaelmeyer/gourd"
)

type stepContext struct {
}

func main() {
  cucumber := &gourd.Cucumber{}

  cucumber.Given(`a scenario with a step`).Do(
    func(context interface{}) {
      _, ok := context.(stepContext)
      if ok {
      }
    })

  cucumber.Given(`no step implemenation`).Pending()

  cucumber.When(`I run cucumber`).Pending()

  cucumber.Then(`cucumber should indicate that the step is undefined`).Pending()

  cucumber.Start();
}

