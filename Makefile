LOCAL_BIN := $(CURDIR)/bin
EASYP_BIN := $(LOCAL_BIN)/easyp
GOIMPORTS_BIN := $(LOCAL_BIN)/goimports
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint
GO_TEST=$(LOCAL_BIN)/gotest
GO_TEST_ARGS=-race -v ./...
UNAME_S := $(shell uname -s)
UNAME_P := $(shell uname -p)
GENERATED_DIR = internal/pb

ARCH :=

ifeq ($(UNAME_S),Linux)
    INSTALL_CMD = sudo apt install -y protobuf-compiler
    ARCH = linux-x86_64
endif

ifeq ($(UNAME_S),Darwin)
    ifeq ($(UNAME_P),arm)
        INSTALL_CMD = brew install protobuf
        ARCH = osx-universal_binary
    else
        INSTALL_CMD = sudo apt install -y protobuf-compiler
        ARCH = linux-x86_64
    endif
endif

all: generate lint test

.PHONY: lint
lint:
	echo 'Running linter on files...'
	$(GOLANGCI_BIN) run \
	--config=.golangci.yaml \
	--max-issues-per-linter=0 \
	--max-same-issues=0

.PHONY: test
test:
	echo 'Running tests...'
	${GO_TEST} ${GO_TEST_ARGS}

.PHONY: run
run:
	go run cmd/main.go

.install-protoc:
	$(INSTALL_CMD)
	which protoc
	protoc --version

bin-deps: .bin-deps

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps: .create-bin
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
	go install github.com/rakyll/gotest@latest && \
	go install github.com/easyp-tech/easyp/cmd/easyp@latest && \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest && \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
	go install golang.org/x/tools/cmd/goimports@latest && \
	go install github.com/envoyproxy/protoc-gen-validate@latest && \
	go install go.uber.org/mock/mockgen@latest && \
	go install github.com/g3co/go-swagger-merger@latest

.create-bin:
	rm -rf ./bin
	mkdir -p ./bin

generate: bin-deps .generate
fast-generate: .generate

.generate:
	$(info Generating code...)

	rm -rf ./$(GENERATED_DIR)
	mkdir ./$(GENERATED_DIR)

	rm -rf ./docs/spec
	mkdir -p ./docs/spec

	rm -rf ~/.easyp/

	(PATH="$(PATH):$(LOCAL_BIN)" && $(EASYP_BIN) mod download && $(EASYP_BIN) generate)

	$(GOIMPORTS_BIN) -w .

	find . -iname "*.swagger.json" -print0 | xargs -0 $(LOCAL_BIN)/go-swagger-merger -o docs/swagger.json

enum: .enum

.enum:
	go install github.com/abice/go-enum@latest
	go generate ./...
