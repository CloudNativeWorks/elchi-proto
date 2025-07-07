# Elchi Proto

This repository contains Protocol Buffers (protobuf) definitions for client-server communication. These definitions are designed to be used across different projects.

## Usage

### In a Go Project

You can add this repository as a Go module to your projects:

```bash
go get github.com/CloudNativeWorks/elchi-proto
```

Then import it in your code:

```go
import "github.com/CloudNativeWorks/elchi-proto"
```

### As a Git Submodule

```bash
git submodule add https://github.com/CloudNativeWorks/elchi-proto.git proto
git submodule update --init --recursive
```

### Generating Go Code

After cloning the repository, you can generate Go code from the proto files:

```bash
make generate
```

## Contents

The repository contains the following proto definitions:
- `client.proto`: Main service definitions
- `commands.proto`: Command types
- `identity.proto`: Authentication
- `register.proto`: Registration processes
- `request.proto`: Request messages
- `response.proto`: Response messages

## Requirements

- Protocol Buffers Compiler (protoc) 3.0+
- Go protocol buffers plugins:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` 