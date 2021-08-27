.PHONY: build run

build:
	go build -v -o bin/go-mock-server ./cmd

run:
	go run ./cmd

clean:
	rm -rf ./bin

.DEFAULT_GOAL := build