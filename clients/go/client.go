package main

import (
	"context"
	"log"

	"github.com/daheige/goapp/clients/go/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	// address     = "localhost:50050" // nginx grpc_pass port
)

/**
% go run client.go
2020/07/23 22:22:41 name:username: xiaoming,message:call ok
*/

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterServiceClient(conn)

	// Contact the server and print out its response.
	res, err := c.SayHello(context.Background(), &pb.HelloReq{
		Id: 1,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("name:%s,message:%s", res.Name, res.Message)
}
