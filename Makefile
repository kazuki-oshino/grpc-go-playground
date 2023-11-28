SHELL=/bin/bash

.PHONY: gen-proto
gen-proto:
	rm -rf ./gen/pb
	mkdir -p ./gen/pb
	protoc --go_out=./gen/pb --go_opt=paths=source_relative \
            --go-grpc_out=./gen/pb --go-grpc_opt=paths=source_relative \
            ./proto/hello.proto
