# gourd

[![Build Status](https://travis-ci.org/raphaelmeyer/gourd.svg?branch=master)](https://travis-ci.org/raphaelmeyer/gourd)
[![Coverage Status](https://img.shields.io/coveralls/raphaelmeyer/gourd.svg)](https://coveralls.io/r/raphaelmeyer/gourd?branch=master)
[![GoDoc](https://godoc.org/github.com/raphaelmeyer/gourd?status.svg)](http://godoc.org/github.com/raphaelmeyer/gourd)

go wire server

## Dependencies

[testify](http://github.com/stretchr/testify)

## Development

### BDD cycle

Run acceptance tests:
```
go run features/wire_server.go
cucumber --wip -t @wip
```

### TDD cycle

Run unit tests:
```
go test
```

Format code:
```
go fmt ./...
```

