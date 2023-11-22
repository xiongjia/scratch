# makefile for stray

CMD_MAIN=stray-main
CMD_CORE=stray-core

ifeq ($(OS),Windows_NT)
	CMD_MAIN=stray-main.exe
	CMD_CORE=stray-core.exe
endif

build: gen
	go build -o ${CMD_MAIN} ./cmd/main/main.go
	go build -o ${CMD_CORE} ./cmd/core/main.go

gen:
	rm -rf ./gen
	$(MAKE) -C api

clean:
	rm -f ${CMD_MAIN} ${CMD_CORE}

.PHONY: build clean
