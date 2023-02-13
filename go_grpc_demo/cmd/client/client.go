package main

import (
	"context"
	v1 "go_grpc/api/proto/v1"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	address := "127.0.0.1:8099"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("server couldn't connect: ", err)
	}
	defer conn.Close()

	c := v1.NewToDoServiceClient(conn)
	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cannel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)

	// create
	req1 := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("create failed %v", err)
	}
	log.Printf("create result: %v", res1)
	id := res1.Id

	// read
	req2 := v1.ReadRequest{Api: apiVersion, Id: id}
	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed %v", err)
	}
	log.Printf("Read result %v", res2)

}
