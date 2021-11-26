# PingPong Server
this is grpc server for testing pbkit packages.

## How to start?
```sh
brew tap riiid/riiid
brew install pbkit
brew install go protoc-gen-go protoc-gen-go-grpc
pollapo install
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I=.pollapo .pollapo/pbkit/interface-pingpong-server/pingpong.proto
go run main.go
```
