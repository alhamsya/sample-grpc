package main

import (
	"context"
	"log"
	"net"

	pb "github.com/alhamsya/sample-grpc/protos"
	"google.golang.org/grpc"
)

type HelloServiceServer struct{}

func main() {
	srv := grpc.NewServer()

	var helloSrv HelloServiceServer

	pb.RegisterHelloServiceServer(srv, helloSrv)

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	errServe := srv.Serve(l)
	if errServe != nil {
		log.Fatal(errServe)
	}
}

func (HelloServiceServer) SayHello(ctx context.Context, pbReq *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Message from client:", pbReq.Greeting)

	return &pb.HelloResponse{
		Replay: "Hello from server",
	}, nil

}
