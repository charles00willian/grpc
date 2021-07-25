package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)
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

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Jao",
			Email: "jao0@gmail.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Jao1",
			Email: "jao1@gmail.com",
		},
		&pb.User{
			Id:    "02",
			Name:  "Jao2",
			Email: "jao02@gmail.com",
		},
		&pb.User{
			Id:    "03",
			Name:  "Jao3",
			Email: "jao03@gmail.com",
		},
		&pb.User{
			Id:    "04",
			Name:  "Jao5",
			Email: "jao05@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
