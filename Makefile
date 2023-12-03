run: 
	@go run main.go

.PHONY: run

build:
	@go build -o ./bin/nobody main.go

.PHONY: build

test:
	@go test -v ./...

.PHONY: test

tidy:
	@go mod tidy

.PHONY: tidy

all: run

.PHONY: all

