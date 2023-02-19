package server

import (
	"context"
	"fmt"
	v1 "go_grpc/api/service/v1"
	"go_grpc/server"
	"go_grpc/storage"
)

type Config struct {
	GRPCPort            string
	DataStoreDBHost     string
	DataStoreDBUser     string
	DataStoreDBPassword string
	DataStoreDBSchema   string
}

func RunServer(port string) error {
	ctx := context.Background()
	//var cfg Config

	// flag.StringVar
	db, err := storage.InitTables()
	if err != nil {
		return fmt.Errorf("create database error: %v", err)
	}
	v1API := v1.NewToDoServiceServer(db)
	return server.RunServer(ctx, v1API, port)
}
