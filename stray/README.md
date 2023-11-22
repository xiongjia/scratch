# tools

go install github.com/jteeuwen/go-bindata/go-bindata

go install github.com/bufbuild/buf/cmd/buf
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Buf

## Buf generate

-   cd api
-   buf generate

## Swagger

-   ui: go-bindata --nocompress -pkg ui -o generated/swagger/ui/ui.go -prefix third_party third_party/swagger-ui/...
-   data: go-bindata --nocompress -pkg data -o generated/swagger/data/data.go -prefix generated/api/swagger generated/api/swagger/...

## Test curl command
- curl http://127.0.0.1:8090/v1/example/echo -d "{ \"Name\": \"test1\" }"
