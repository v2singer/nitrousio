package main

import (
	"fmt"
	"go_grpc/cmd/server"
)

func main() {
	port := "8099"
	err := server.RunServer(port)
	fmt.Println(err.Error())
}
