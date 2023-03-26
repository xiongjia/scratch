# makefile for stray

COMMANDS = stray 

build: $(COMMANDS)

clean:
	go clean
	rm -f ${COMMANDS}

stray: $(shell find ./cmd/stray -iname "*.go") $(shell find ./util -iname "*.go") 
	go build -o stray ./cmd/stray/main.go

.PHONY: build clean
