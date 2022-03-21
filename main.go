package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	interface_pingpong_server "pbkit/pingpong-server/pbkit/interface-pingpong-server"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	httpPort = "8080"
)

type PingPongServer struct {
	interface_pingpong_server.UnimplementedPingPongServiceServer
}

type ThrowServiceServer struct {
	interface_pingpong_server.UnimplementedThrowServiceServer
}

func (s *PingPongServer) PingPong(ctx context.Context, in *interface_pingpong_server.Ping) (*interface_pingpong_server.Pong, error) {
	log.Printf("Received: %s", in.GetHello())
	return &interface_pingpong_server.Pong{
		World: "Pong",
	}, nil
}

func (s *ThrowServiceServer) Throw(ctx context.Context, in *interface_pingpong_server.ThrowRequest) (*interface_pingpong_server.ThrowResponse, error) {
	log.Printf("Received: %s", in.GetCode().Enum())
	return nil, status.Errorf(codes.Code(in.GetCode().Number()), "%s", in.GetCode().Enum())
}

func main() {
	grpcServer := grpc.NewServer()
	interface_pingpong_server.RegisterPingPongServiceServer(grpcServer, &PingPongServer{})
	interface_pingpong_server.RegisterThrowServiceServer(grpcServer, &ThrowServiceServer{})

	wrappedServer := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		if wrappedServer.IsGrpcWebRequest(req) || wrappedServer.IsAcceptableGrpcCorsRequest(req) {
			wrappedServer.ServeHTTP(resp, req)
			return
		}
		http.DefaultServeMux.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%v", httpPort),
		Handler: http.HandlerFunc(handler),
	}

	log.Printf("gRPC web server listening at %v", httpPort)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to start http server at %v", httpPort)
	}
}
