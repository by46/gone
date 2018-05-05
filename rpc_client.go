package main

import (
	"context"
	"log"
	"time"

	"github.com/by46/gone/im"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := im.NewGreeterClient(conn)

	name := "benjamin"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &im.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r, err = client.SayHelloAgain(ctx, &im.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
