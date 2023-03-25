# makefile for stray

COMMANDS = stray 

build: $(COMMANDS)

clean:
	go clean
	rm -f ${COMMANDS}

stray:
	go build -o stray ./cmd/stray/main.go

.PHONY: build clean
