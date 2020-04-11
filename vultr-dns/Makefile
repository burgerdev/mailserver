
all: main

main: fmt main.go
	go build -o main ./main.go

fmt:
	gofmt -w ./*.go

run: main
	./main

vendor:
	go mod vendor

.PHONY: all run fmt
