.PHONY: generate clean tools

PROTO_FILES=$(shell find client -name "*.proto")
GO_OUT_DIR=.

# Pinned codegen toolchain — keep protoc-gen-go aligned with the
# google.golang.org/protobuf runtime in go.mod so generated code and runtime
# never drift. Run `make tools` once (then `make generate`).
PROTOC_GEN_GO_VERSION ?= v1.36.6
PROTOC_GEN_GO_GRPC_VERSION ?= v1.5.1

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

generate:
	protoc -I. \
		--go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

clean:
	find . -name "*.pb.go" -type f -delete
