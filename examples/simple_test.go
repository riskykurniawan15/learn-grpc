package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/riskykurniawan15/learn-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	fmt.Println("🚀 Testing User Service...")

	// Create user
	fmt.Println("1️⃣ Creating user...")
	user, err := client.CreateUser(ctx, &proto.CreateUserRequest{
		Name:  "Alice Johnson",
		Email: "alice@example.com",
		Age:   28,
	})
	if err != nil {
		log.Printf("❌ Create failed: %v", err)
		return
	}
	fmt.Printf("✅ User created: %s (%s)\n", user.User.Name, user.User.Email)

	// Get user
	fmt.Println("\n2️⃣ Getting user...")
	retrieved, err := client.GetUser(ctx, &proto.GetUserRequest{Id: user.User.Id})
	if err != nil {
		log.Printf("❌ Get failed: %v", err)
		return
	}
	fmt.Printf("✅ User retrieved: %s, Age: %d\n", retrieved.User.Name, retrieved.User.Age)

	// Update user
	fmt.Println("\n3️⃣ Updating user...")
	updated, err := client.UpdateUser(ctx, &proto.UpdateUserRequest{
		Id:   user.User.Id,
		Age:  29,
		Name: "Alice Johnson Updated",
	})
	if err != nil {
		log.Printf("❌ Update failed: %v", err)
		return
	}
	fmt.Printf("✅ User updated: %s, Age: %d\n", updated.User.Name, updated.User.Age)

	// Get all users
	fmt.Println("\n4️⃣ Getting all users...")
	allUsers, err := client.GetAllUsers(ctx, &proto.GetAllUsersRequest{})
	if err != nil {
		log.Printf("❌ Get all failed: %v", err)
		return
	}
	fmt.Printf("✅ Found %d users\n", len(allUsers.Users))

	// Delete user
	fmt.Println("\n5️⃣ Deleting user...")
	deleted, err := client.DeleteUser(ctx, &proto.DeleteUserRequest{Id: user.User.Id})
	if err != nil {
		log.Printf("❌ Delete failed: %v", err)
		return
	}
	fmt.Printf("✅ %s\n", deleted.Message)

	fmt.Println("\n🎉 All tests completed successfully!")
}
