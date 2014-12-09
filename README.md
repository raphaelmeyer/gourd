# gourd

[![Build Status](https://travis-ci.org/raphaelmeyer/gourd.png)](https://travis-ci.org/raphaelmeyer/gourd)
[![Coverage Status](https://coveralls.io/repos/raphaelmeyer/gourd/badge.png)](https://coveralls.io/r/raphaelmeyer/gourd)
[![GoDoc](https://godoc.org/github.com/raphaelmeyer/gourd?status.png)](http://godoc.org/github.com/raphaelmeyer/gourd)

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

