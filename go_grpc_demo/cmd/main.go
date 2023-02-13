package main

import (
	"fmt"
	"go_grpc/cmd/server"
)

func main() {
	err := server.RunServer()
	fmt.Println(err.Error())
}
