# makefile for stray

CMD_MAIN=stray-infra
CMD_CORE=stray-core

ifeq ($(OS),Windows_NT)
	CMD_MAIN=stray-infra.exe
	CMD_CORE=stray-core.exe
endif

app:
	go build -o ${CMD_MAIN} ./cmd/infra/main.go
	go build -o ${CMD_CORE} ./cmd/core/main.go

build: generate app

generate:
	rm -rf ./gen
	$(MAKE) -C api
	go-bindata --nocompress -pkg doc -o gen/api_doc/core/v1/doc.go -prefix gen/api_doc/core/v1/ gen/api_doc/core/v1/...
	go-bindata --nocompress -pkg doc -o gen/api_doc/infra/v1/doc.go -prefix gen/api_doc/infra/v1/ gen/api_doc/infra/v1/...
	go-bindata --nocompress -pkg ui -o gen/swagger/ui.go -prefix third_party third_party/swagger-ui/...

clean:
	rm -f ${CMD_MAIN} ${CMD_CORE}

.PHONY: build clean
