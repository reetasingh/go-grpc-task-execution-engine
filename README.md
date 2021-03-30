# run client
```bash
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go start abc
2021/03/29 19:04:12 UUID: 1, Status: RECEIVED, Details: Recieved by Server
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go status 1
2021/03/29 19:04:25 UUID: 1, Status: RECEIVED, Details: Recieved by Server
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ go run cmd/client/main.go stop 1
2021/03/29 19:04:38 UUID: 1, Status: CANCELLED, Details: Task Successfully Cancelled
reetasingh-ltm8:go-grpc-task-execution-engine reetasingh$ 

```

# run server (another terminal)
```bash
reetasingh-ltm8:cmd reetasingh$ go run server/main.go 
2021/03/29 18:57:07 Received Request to Cancel Task: 1
2021/03/29 18:57:23 Received Request for Task Status: 1
2021/03/29 18:58:02 Received Request for Task Status: 1
2021/03/29 18:58:08 Received Request for Task Status: 1
2021/03/29 18:58:24 Received Request for Task Status: 2
2021/03/29 19:03:30 Received Request for Task Status: 1
2021/03/29 19:04:12 Received Request to Execute Task: abc
2021/03/29 19:04:25 Received Request for Task Status: 1
2021/03/29 19:04:38 Received Request to Cancel Task: 1


```

# compile proto files
```bash
reetasingh-ltm8:pkg reetasingh$ protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     task/task.proto
reetasingh-ltm8:pkg reetasingh$ 
```
