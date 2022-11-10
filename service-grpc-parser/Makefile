init:
		go mod tidy
		make compile-grpc
.PHONY: init
build:
		make init
		go build -o build/service-grpc-parser -v ./cmd/app
.PHONY: build
run:
		make build
		./build/service-grpc-parser
.PHONY: run
pretest: ### get linters
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go get -u golang.org/x/lint/golint@latest
.PHONY: pretest
test:
	make pretest
	golint ./...
	go vet ./...
	staticcheck ./...
.PHONY: test
## Compile the proto file.
compile-grpc:
		protoc -I api/proto \
		-I./third_party/grpc-gateway \
        -I./third_party/googleapis \
		--go_out=. \
		--go-grpc_opt=require_unimplemented_servers=false \
		--go-grpc_out=. \
		api/proto/*.proto
.PHONY: compile-grpc
.DEFAULT_GOAL := run