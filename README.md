## A simple application which mimics task execution using GRPC communication between client and server

1. Client sends request to start a task Server and server sends back status of Task
2. Client queries status of Task and server sends back the status
3. Client sends request to cancel Task and server updates task status and returns status of Task


# run server 
```bash
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/server/main.go 
2021/03/30 20:28:45 Received Request to Execute Task: abc
2021/03/30 20:28:55 Received Request to Cancel Task: abc
2021/03/30 20:29:00 Received Request for Task Status: abc
2021/03/30 20:29:05 Received Request for Task Status: abc
2021/03/30 20:29:18 Received Request to Execute Task: pqr
2021/03/30 20:29:33 Received Request for Task Status: pqr
2021/03/30 20:29:37 Received Request for Task Status: abc
2021/03/30 20:31:08 Received Request for Task Status: abc

```


# run client (another terminal)
```bash
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go start abc
2021/03/30 20:28:45 UUID: abc, Name: abc, Status: RECEIVED, Details: Recieved by Server
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go stop abc
2021/03/30 20:28:55 UUID: abc, Name: abc, Status: CANCELLED, Details: Task Successfully Cancelled
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go status abc
2021/03/30 20:29:00 UUID: abc, Name: abc, Status: CANCELLED, Details: Task Successfully Cancelled
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go status abc
2021/03/30 20:29:05 UUID: abc, Name: abc, Status: CANCELLED, Details: Task Successfully Cancelled
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go start pqr
2021/03/30 20:29:18 UUID: pqr, Name: pqr, Status: RECEIVED, Details: Recieved by Server
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go status pqr
2021/03/30 20:29:33 UUID: pqr, Name: pqr, Status: RECEIVED, Details: Recieved by Server
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go status abc
2021/03/30 20:29:37 UUID: abc, Name: abc, Status: CANCELLED, Details: Task Successfully Cancelled
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go status abc
2021/03/30 20:31:08 UUID: abc, Name: abc, Status: CANCELLED, Details: Task Successfully Cancelled
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 


```

# compile proto files
```bash
reetasingh-ltm8:pkg reetasingh$ protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     task/task.proto
reetasingh-ltm8:pkg reetasingh$ 
```
