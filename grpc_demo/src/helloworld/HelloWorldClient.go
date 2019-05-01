package main

import (
	"google.golang.org/grpc"
	"log"
	pb "./proto"
	"os"
	"context"
)

const (
	address = "localhost:50502"
	defaultName = "world"
)


func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", conn)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %v", r.Message)


}
