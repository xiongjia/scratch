# makefile for stray

COMMANDS = stray 
ifeq ($(OS),Windows_NT)
	COMMANDS = stray.exe
endif

LIBS = ./util/conf.go ./util/log.go

build: $(COMMANDS)

clean:
	go clean
	rm -f ${COMMANDS}


$(COMMANDS): ./cmd/stray/main.go $(LIBS)
	go build -o ${COMMANDS} ./cmd/stray/main.go

.PHONY: build clean
