# `go-grpc-pingpong` 🐹🏓

[![developer](https://img.shields.io/badge/developer:-Henrik-blue?style)](https://github.com/henrikalbihn)
![language](https://img.shields.io/badge/language-golang-blue?style)
[![ci](https://github.com/henrikalbihn/go-grpc-pingpong/actions/workflows/build-and-test.yml/badge.svg?query=event=push+branch=master)](https://github.com/henrikalbihn/go-grpc-pingpong/actions/workflows/build-and-test.yml?query=event=push+branch=master)

## Description

A simple gRPC boilerplate project in Golang.

As my first foray into gRPC and Golang, I wanted to create a simple server and client. The server will listen for an empty TCP/gRPC message (`payload={}`) from the client and respond with a `Pong` message.

### Features

- [x] gRPC server
- [x] gRPC client
- [x] gRPC protocol buffer spec
- [x] gRPC code generation (`protoc` and `protoc-gen-go`)
- [x] gRPC server and client binaries
- [x] `make` commands
- [x] github actions CI/CD pipeline

## Usage

### Build

```bash
# to build the binaries 
make build
```

### Run the Server

```bash
# to run the server
make serve
```

```log
2023/09/07 11:03:02 Starting server on port 50051
```

#### Custom Port

```bash
# or with a custom port
./bin/serve -port 50000
```

```log
2023/09/07 11:05:18 Starting server on port 50000
```

### Ping the Server

```bash
# in another terminal
make ping
```

```log
2023/09/07 11:03:04 Establishing gRPC connection - localhost:50051...
2023/09/07 11:03:04 Response: Pong
```

#### Custom Port

```bash
# or with custom port
./bin/ping -port 50000
```

```log
2023/09/07 11:03:04 Establishing gRPC connection - localhost:50000...
2023/09/07 11:03:04 Response: Pong
```

### Stop the Server

```bash
# in the server terminal
# to stop the server, press Ctrl+C
```
<!-- 
## Steps

0. Install Go and add to `${PATH}`.

```bash
yum install golang -y
# or
brew install go@1.20

export PATH=$PATH:$(go env GOPATH)/bin
```

1. Create a new directory for the project and change into it.

```bash
mkdir go-grpc-pingpong
cd go-grpc-pingpong
```

2. Create a new module for the project.

```bash
go mod init github.com/henrikalbihn/go-grpc-pingpong
```

3. Add dependencies to the project.

```bash
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go
```

4. Create protocol buffer spec `pingpong.proto`.

5. Generate the gRPC code.

```bash
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  pingpong.proto
``` -->

## References

- [gRPC Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [gRPC Go Tutorial](https://grpc.io/docs/languages/go/basics/)
- [Protobuf-Go Generated Code Guide](https://protobuf.dev/reference/go/go-generated/)
- [grpc-go Hello World Example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)
