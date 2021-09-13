package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	interface_pingpong_server "pbkit/pingpong-server/pbkit/interface-pingpong-server"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

const (
	httpPort = "8080"
)

type PingPongServer struct {
	interface_pingpong_server.UnimplementedPingPongServiceServer
}

func (s *PingPongServer) PingPong(ctx context.Context, in *interface_pingpong_server.Ping) (*interface_pingpong_server.Pong, error) {
	log.Printf("Received: %s", in.GetHello())
	return &interface_pingpong_server.Pong{
		World: "Pong",
	}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	interface_pingpong_server.RegisterPingPongServiceServer(grpcServer, &PingPongServer{})

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		if wrappedServer.IsGrpcWebRequest(req) {
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
