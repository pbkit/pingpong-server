pre-install: 
	brew install protoc-gen-go protoc-gen-go-grpc

gen: 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I=.pollapo -I=$(shell pb vendor directory) .pollapo/pbkit/interface-pingpong-server/pingpong.proto

dev:
	go run main.go