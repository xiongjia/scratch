# Buf 

## Buf generate
- cd api
- buf generate

## Swagger 
- ui  go-bindata --nocompress -pkg ui -o generated/swagger/ui/ui.go -prefix third_party third_party/swagger-ui/...
- data  go-bindata --nocompress -pkg data -o generated/swagger/data/data.go -prefix generated/api/swagger generated/api/swagger/...
	
	
