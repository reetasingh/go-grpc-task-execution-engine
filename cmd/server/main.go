
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

type TaskExecutionServer struct {
	pb.UnimplementedTaskExecutionServer
	taskMap map[string]*pb.TaskResponse
}

type TaskNotFoundError struct {
    key string
}

func (t *TaskNotFoundError) Error() string {
    return "Key not found in map"
}

func NewTaskNotFoundError(key string) *TaskNotFoundError {
    var t TaskNotFoundError
    t.key = key
    return &t
}


func NewTaskExecutionServer() *TaskExecutionServer {
    var s TaskExecutionServer
    s.taskMap = make(map[string]*pb.TaskResponse)
    return &s
}

func (s *TaskExecutionServer) ExecuteTask(ctx context.Context, in *pb.TaskExecutionRequest) (*pb.TaskResponse, error) {
	log.Printf("Received Request to Execute Task: %v", in.GetName())

	taskResponse := pb.TaskResponse{Uuid: in.GetName(), Status: "RECEIVED", Details: "Recieved by Server", Name: in.GetName()}
	s.taskMap[in.GetName()] = &taskResponse
	return &taskResponse, nil
}

func (s *TaskExecutionServer) GetTaskStatus(ctx context.Context, in *pb.TaskStatusRequest) (*pb.TaskResponse, error) {
	log.Printf("Received Request for Task Status: %v", in.GetUuid())

	if val, ok := s.taskMap[in.GetUuid()]; ok {
	    return val, nil
    }
    err := NewTaskNotFoundError(in.GetUuid())

    return nil, err
}

func (s *TaskExecutionServer) CancelTask(ctx context.Context, in *pb.TaskStatusRequest) (*pb.TaskResponse, error) {
	log.Printf("Received Request to Cancel Task: %v", in.GetUuid())
	if val, ok := s.taskMap[in.GetUuid()]; ok {
	    val.Status = "CANCELLED"
	    val.Details = "Task Successfully Cancelled"
	    return val, nil
    }

	err := NewTaskNotFoundError(in.GetUuid())
    return nil, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	newServer := NewTaskExecutionServer()
	pb.RegisterTaskExecutionServer(s, newServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}