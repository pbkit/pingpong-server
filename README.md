# PingPong Server
this is grpc server for testing pbkit packages.

## Before Start

Require installing [pollapo](https://pbkit.dev/docs/getting-started/installation), [go](https://go.dev/doc/install), [protoc](https://github.com/protocolbuffers/protobuf/releases/tag/v21.8).

(Require installing [nodejs](https://nodejs.org/ko/download/) and [yarn](https://yarnpkg.com/getting-started/install) for client.)

And Install `protoc-gen-go` and `protoc-gen-go-grpc`.

- MacOS
    - Run `brew install protoc-gen-go protoc-gen-go-grpc`
- Windows
    - Run `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
    - Run `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

Clone pingpong-server and api-common-protos.

```sh
git clone https://github.com/pbkit/pingpong-server.git
git clone https://github.com/googleapis/api-common-protos.git
```

## How to start?

Gen go file for protoc.

```sh
cd ./pingpong-server

pollapo install

protoc --go_out="." --go_opt=paths=source_relative --go-grpc_out="." --go-grpc_opt=paths=source_relative -I=".pollapo" .pollapo/pbkit/interface-pingpong-server/pingpong.proto -I="../api-common-protos/"
```

And start pingpong server.

```sh
go run main.go
```

### Then...

Clone [pingpong-client](https://github.com/pbkit/pingpong-client) and Install [pbkit-devtools](https://github.com/pbkit/pbkit-devtools).

And Enjoy Coding!