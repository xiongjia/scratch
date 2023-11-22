# makefile for stray

COMMANDS = stray
ifeq ($(OS),Windows_NT)
    COMMANDS = stray.exe
endif

build: $(COMMANDS)

clean:
	rm -f ${COMMANDS}


$(COMMANDS): ./cmd/stray/main.go
	go build -o ${COMMANDS} ./cmd/stray/main.go

.PHONY: build clean
