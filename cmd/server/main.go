
// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "go-grpc-task-execution-engine/pkg/helloworld"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedTaskExecutionServer
}

func (s *server) ExecuteTask(ctx context.Context, in *pb.TaskExecutionRequest) (*pb.TaskResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.TaskResponse{Uuid: "1", Status: "RECEIVED", Details: "Recieved by Server"}, nil
}

func (s *server) GetTaskStatus(ctx context.Context, in *pb.TaskStatusRequest) (*pb.TaskResponse, error) {
	log.Printf("Received: %v", in.GetUuid())
	return &pb.TaskResponse{Uuid: "1", Status: "RECEIVED", Details: "Recieved by Server"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskExecutionServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}