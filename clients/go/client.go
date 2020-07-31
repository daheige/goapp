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

	res2, err := c.Info(context.Background(), &pb.InfoReq{
		Name: "daheige",
	})

	log.Println(res2, err)
}

/**
2020/07/31 23:08:14 name:username: xiaoming,message:call ok
2020/07/31 23:08:14 address:"shenzhen" message:"ok" <nil>

当参数不合法抛出错误
2020/07/31 23:07:28 name:username: xiaoming,message:call ok
2020/07/31 23:07:28 <nil> rpc error: code = InvalidArgument desc = Key: 'InfoReq.Name' Error:Field validation
for 'Name' failed on the 'required' tag
*/
