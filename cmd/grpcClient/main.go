package main

import (
	"context"
	"log"
	"time"

	pb "github.com/gabogomes/Go-gRPC-Example/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client from the generated code.
	client := pb.NewCategoryServiceClient(conn)

	// Prepare the request you want to send to the server.
	req := &pb.CategoryGetRequest{Id: "98c81582-bdc1-46fe-b8b7-ad59a10de10b"}

	// Send the request to the server and get a response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetCategory(ctx, req)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	// Output the response.
	log.Printf("Response from server: %v", res)
}
