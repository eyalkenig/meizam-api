# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: build

# Omit the debug information
LD_FLAGS := -w

default: gofmt dependencies test build

gofmt:
	./build/gofmt.sh

gofmt-w:
	./build/gofmt.sh fix

getdep:
	go get github.com/golang/dep/cmd/dep

dependencies: getdep
	dep ensure

getgotest:
	# Like go test but with colors
	go get github.com/rakyll/gotest

test: getgotest
	gotest -parallel 1 -v ./app/...

build:
	go build -ldflags "$(LD_FLAGS)" -o build/out/go-meizam-api main.go

system_tests: getgotest
	gotest -parallel 1 -v ./system_tests/...
