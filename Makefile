TEST?=$$(go list ./... | grep -v 'vendor')
BINARY=go-grpc-task-execution-engine

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

default: build

build:
	go build cmd -o ${BINARY}

fmt:
	@gofmt -l -w $(SRC)

lint:
	golangci-lint run

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
	xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

validate-modules:
	@echo "- Verifying that the dependencies have expected content..."
	go mod verify
	@echo "- Checking for any unused/missing packages in go.mod..."
	go mod tidy
	@echo "- Checking for unused packages in vendor..."
	go mod vendor
	@git diff --exit-code -- go.sum go.mod vendor/