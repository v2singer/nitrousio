package main

import (
	"context"
	"fmt"
	v1 "go_grpc/api/proto/v1"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func create(ctx context.Context, c v1.ToDoServiceClient, stopCreate chan struct{}, ids chan int64) {
	for {
		t := time.Now().In(time.UTC)
		reminder, _ := ptypes.TimestampProto(t)
		pfx := t.Format(time.RFC3339Nano)

		select {
		case <-stopCreate:
			close(ids)
			return
		case <-time.After(time.Second * 1):
		}

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
		ids <- res1.Id

	}
}

func read(ctx context.Context, c v1.ToDoServiceClient, ids chan int64, done chan struct{}) {

	for {
		select {
		case id, ok := <-ids:
			// closed
			if !ok {
				done <- struct{}{}
				return
			}
			// read
			req2 := v1.ReadRequest{Api: apiVersion, Id: id}
			res2, err := c.Read(ctx, &req2)
			if err != nil {
				log.Fatalf("Read failed %v", err)
			}
			log.Printf("Read result %v", res2)
		}
	}
}

func stop(s chan struct{}) {
	s <- struct{}{}
}

func waitDone(d chan struct{}) {
	<-d
}

func main() {
	address := "127.0.0.1:8099"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("server couldn't connect: ", err)
	}
	defer conn.Close()

	c := v1.NewToDoServiceClient(conn)
	ctx, cannel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cannel()

	ids := make(chan int64, 10)
	done := make(chan struct{})
	stopCreate := make(chan struct{})

	go create(ctx, c, stopCreate, ids)

	go read(ctx, c, ids, done)

	time.Sleep(time.Second * 10)
	stop(stopCreate)
	waitDone(done)
	fmt.Println("done")
}
