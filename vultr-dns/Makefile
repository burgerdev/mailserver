all: fmt main

clean:
	rm -rf static main vendor/

main: vendor main.go
	go build -o main ./main.go

static: vendor main.go
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o static ./main.go

fmt:
	gofmt -w ./*.go

run: main
	./main

vendor:
	go mod vendor

.PHONY: all run fmt
