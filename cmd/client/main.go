package main

import (
	"context"
	"fmt"
	pb "github.com/kazuki-oshino/grpc-go-playground/gen/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	fmt.Println("start gRPC Client.")

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx := context.Background()
	res, err := client.SayHello(
		ctx,
		&pb.HelloRequest{
			Name:  "gRPC Client",
			Age:   10,
			Hello: "helllllllooooooooooooo",
		},
	)
	if err != nil {
		log.Fatal("Failed to say hello.", err)
		return
	}

	fmt.Println(res)
}
