// Package main implements a client for Greeter service.
package main

import (
	"go-grpc-task-execution-engine/pkg/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
    helloworld.Hello()
}