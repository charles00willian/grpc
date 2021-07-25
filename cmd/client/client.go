package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gihub.com/charles00willian/grcp-go-starter/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	// stablish connection to server
	client := pb.NewUserServiceClient(connection)
	AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatal("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Joao",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatal("Could not make gRPC request: %v", err)
	}

	// loop which verifies the streaming
	for {
		// starts receiving info
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Could not receive the msg: %v", err)
		}

		fmt.Println("Status:", stream.Status)
	}
}
