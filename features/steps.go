package main

import (
  "net"
  "github.com/raphaelmeyer/gourd"
)

type stepContext struct {
  testee gourd.Cucumber
}

func main() {
  cucumber := &gourd.Cucumber{}

  cucumber.Given(`a step with no implementation`).Do(
    func(context interface{}) {
      _, ok := context.(stepContext)
      if ok {
      }
    })

  cucumber.When(`I run cucumber`).Do(
    func(context interface{}) {
      stepContext, ok := context.(stepContext)
      if ok {
        stepContext.testee.Start()
        conn, err := net.Dial("tcp", "localhost:1847")
        cucumber.Assert(err == nil)
        conn.Close()
      }
    })

  cucumber.Then(`cucumber should indicate that the step is undefined`).Pending()

  cucumber.Start();
}

