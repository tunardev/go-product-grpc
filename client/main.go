package main

import (
	"context"
	"fmt"

	pb "github.com/tunardev/go-product-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	// Create a gRPC connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewProductServiceClient(conn)
	req := &pb.Product{
		Id: 132,
		Name: "Vision Pro",
		Price: 3500,
	}

	// Call the service
	res, err := client.GetProduct(context.Background(), req) 
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: ID=%d, Name=%s, Price=%.2f, Brand=%s\n", res.Id, res.Name, res.Price, res.Brand)
}
