package main

import (
	"log"
	"net"

	"github.com/riskykurniawan15/learn-grpc/database"
	"github.com/riskykurniawan15/learn-grpc/proto"
	"github.com/riskykurniawan15/learn-grpc/service"
	"github.com/riskykurniawan15/learn-grpc/validation"
	"google.golang.org/grpc"
)

func main() {
	// Initialize database
	database.InitDatabase()

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register user service
	userService := service.NewUserService(validation.NewValidator())
	proto.RegisterUserServiceServer(grpcServer, userService)

	// Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server starting on port 50051...")

	// Start serving
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
