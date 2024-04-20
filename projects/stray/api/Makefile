# makefile for stray

COMMANDS = stray
ifeq ($(OS),Windows_NT)
    COMMANDS = stray.exe
endif

all: build

build:
	buf generate

.PHONY: build
