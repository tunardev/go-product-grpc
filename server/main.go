package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/tunardev/go-product-grpc/pb"
	"google.golang.org/grpc"
)

type productServiceServer struct{
	pb.UnimplementedProductServiceServer
}

func (s *productServiceServer) GetProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	product := &pb.Product{
		Id:    req.Id,
		Name:  req.Name,
		Price: req.Price,
	}

	if product.Name == "Vision Pro" {
		product.Brand = "Apple"
	} else {
		product.Brand = "Samsung"
	}
	return product, nil
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer() // Create a gRPC server

	// // Register the service implementation to the gRPC server
	pb.RegisterProductServiceServer(grpcServer, &productServiceServer{})
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	fmt.Println("gRPC server started on port 50051")
}
