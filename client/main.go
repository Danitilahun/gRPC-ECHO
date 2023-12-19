package main

import (
	"context"
	"fmt"

	hello "github.com/Danitilahun/gRPC-Hello_World.git/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := hello.NewGreeterClient(conn)
	resp, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "Bilal"})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Message)
}
