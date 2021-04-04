package main

import (
	"context"
	"log"

	pb "github.com/alhamsya/sample-grpc/protos"
	"google.golang.org/grpc"
)

func main() {
	req := pb.HelloRequest{
		Greeting: "hello",
	}

	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloServiceClient(conn)

	resp, err := client.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
