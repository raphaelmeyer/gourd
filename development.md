# gourd development

## Setup

Install git, go, cucumber

Clone gourd into your GOPATH:
```
mkdir -p $GOPATH/src/github.com/raphaelmeyer/
cd $GOPATH/src/github.com/raphaelmeyer/

git clone https://github.com/raphaelmeyer/gourd.git
```

Get the dependencies:
```
cd $GOPATH/src/github.com/raphaelmeyer/gourd

go get -t ./...
```

Test:
```
cd $GOPATH/src/github.com/raphaelmeyer/gourd

go test
```

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

