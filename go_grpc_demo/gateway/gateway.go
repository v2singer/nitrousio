package gateway

import (
	"context"
	"fmt"
	v1 "go_grpc/api/proto/v1"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run http listen
func Run() error {
	grpcAddr := "127.0.0.1:8099"
	httpAddr := "127.0.0.1:8199"

	ctx := context.Background()
	op := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.DialContext(ctx, grpcAddr, op...)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return err
	}

	// customerMd := func(ctx context.Context, request *http.Request) metadata.MD {
	// userID := request.Header.Get("x-user-id")
	// md := make(map[string]string)
	// md["x-user-id"] = userID
	// return metadata.New(md)
	// }

	//mux := runtime.NewServeMux(runtime.WithMetadata(customerMd))
	mux := runtime.NewServeMux()

	err = newGateway(ctx, conn, mux)
	if err != nil {
		fmt.Printf("register hander failed: %v\n", err)
		return err
	}

	server := http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}
	fmt.Printf("Serving Http on %s\n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func newGateway(ctx context.Context, conn *grpc.ClientConn, mux *runtime.ServeMux) error {
	err := v1.RegisterToDoServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}
	return nil
}
