package server

import (
	"context"
	v1 "go_grpc/api/proto/v1"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)
	c := make(chan os.Signal, 1)
	go func() {
		for range c {
			log.Println("shutting down gRPC server....")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	log.Println("starting gRPC server....")
	return server.Serve(listen)
}
