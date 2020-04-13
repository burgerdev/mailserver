
include secrets/env

SOURCES=main.go lib.go

all: fmt main

clean:
	rm -rf static main vendor/

main: vendor $(SOURCES)
	go build -o main ./

static: vendor $(SOURCES)
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o static ./

fmt:
	gofmt -w ./*.go

run: fmt main
	./main

vendor:
	go mod vendor

.PHONY: all run fmt
