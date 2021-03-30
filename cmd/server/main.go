
// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "go-grpc-task-execution-engine/pkg/task"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedTaskExecutionServer
}

func (s *server) ExecuteTask(ctx context.Context, in *pb.TaskExecutionRequest) (*pb.TaskResponse, error) {
	log.Printf("Received Request to Execute Task: %v", in.GetName())
	taskResponse := pb.TaskResponse{Uuid: "1", Status: "RECEIVED", Details: "Recieved by Server", Name: in.GetName()}
	return &taskResponse, nil
}

func (s *server) GetTaskStatus(ctx context.Context, in *pb.TaskStatusRequest) (*pb.TaskResponse, error) {
	log.Printf("Received Request for Task Status: %v", in.GetUuid())
	return &pb.TaskResponse{Uuid: in.GetUuid(), Status: "RECEIVED", Details: "Recieved by Server", Name: "change me"}, nil
}

func (s *server) CancelTask(ctx context.Context, in *pb.TaskStatusRequest) (*pb.TaskResponse, error) {
	log.Printf("Received Request to Cancel Task: %v", in.GetUuid())
	return &pb.TaskResponse{Uuid: in.GetUuid(), Status: "CANCELLED", Details: "Task Successfully Cancelled", Name: "change me"}, nil
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