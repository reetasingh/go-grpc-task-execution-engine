
// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "go-grpc-task-execution-engine/pkg/task"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTaskExecutionClient(conn)

	action := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if action == "start" {
	    name := os.Args[2]
	    r, err := c.ExecuteTask(ctx, &pb.TaskExecutionRequest{Name: name})
	    if err != nil {
		    log.Fatalf("could not start: %v", err)
	    }
	    log.Printf("UUID: %s, Name: %s, Status: %s, Details: %s", r.GetUuid(), r.GetName(), r.GetStatus(), r.GetDetails())
	} else if action == "stop" {

	    uuid := os.Args[2]
	    r, err := c.CancelTask(ctx, &pb.TaskStatusRequest{Uuid: uuid})
	    if err != nil {
		    log.Fatalf("could not cancel: %v", err)
	    }
	    log.Printf("UUID: %s, Name: %s, Status: %s, Details: %s", r.GetUuid(), r.GetName(), r.GetStatus(), r.GetDetails())
	} else if action == "status" {
	    uuid := os.Args[2]
	    r, err := c.GetTaskStatus(ctx, &pb.TaskStatusRequest{Uuid: uuid})
	    if err != nil {
		    log.Fatalf("could not get status: %v", err)
	    }
	    log.Printf("UUID: %s, Name: %s, Status: %s, Details: %s", r.GetUuid(), r.GetName(), r.GetStatus(), r.GetDetails())
	}

}