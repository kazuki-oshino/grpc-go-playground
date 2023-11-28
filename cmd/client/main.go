package main

import (
	"context"
	"fmt"
	pb "github.com/kazuki-oshino/grpc-go-playground/gen/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func main() {
	fmt.Println("start gRPC Client.")

	address := "localhost:8080"
	keepAliveClientParam := keepalive.ClientParameters{
		Time:                80 * time.Second,
		Timeout:             5 * time.Second,
		PermitWithoutStream: true,
	}
	fmt.Println("start dial.")
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepAliveClientParam),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()
	fmt.Println("end dial.")

	time.Sleep(60 * time.Second)
	client := pb.NewGreeterClient(conn)
	ctx := context.Background()
	res, err := client.SayHello(
		ctx,
		&pb.HelloRequest{
			Name:  "bob",
			Age:   10,
			Hello: "helllllllooooooooooooo",
		},
	)
	if err != nil {
		log.Fatal("Failed to say hello.", err)
		return
	}

	fmt.Println(res)
	time.Sleep(20 * time.Second)

	res, err = client.SayHello(
		ctx,
		&pb.HelloRequest{
			Name: "bob2",
		},
	)
	if err != nil {
		log.Fatal("Failed to say hello.", err)
		return
	}

	fmt.Println(res)
}
