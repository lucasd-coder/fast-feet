
FIND ?= $(shell find proto/ -iname "*.proto")

GOPATH ?= go env GOPATH

.PHONY: proto-gen generate-mocks wire-gen run-application test coverage

proto-gen:	
	protoc --proto_path=proto/ $(FIND) \
		--plugin=$(GOPATH)/bin/protoc-gen-go-grpc \
		--go-grpc_out=. --go_out=.;

generate-mocks:
	./scripts/mockery.sh


wire-gen:
	cd internal/app && wire

run-application:
	GO_PROFILE=dev GO111MODULE=on go run ./cmd/app/main.go

test:
	GO111MODULE=on go test -coverprofile coverage.out ./...

coverage:
	GO111MODULE=on go test -coverprofile coverage.out ./...
	GO111MODULE=on go tool cover -html=coverage.out

