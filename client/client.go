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
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create client
	client := proto.NewUserServiceClient(conn)

	// Set timeout context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	fmt.Println("=== Go RPC User Service Test ===")

	// Test 1: Create User
	fmt.Println("1. Creating user...")
	createResp, err := client.CreateUser(ctx, &proto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
		Age:      30,
	})
	if err != nil {
		log.Printf("CreateUser failed: %v", err)
	} else {
		fmt.Printf("   User created: ID=%d, Name=%s, Email=%s, Age=%d\n",
			createResp.User.Id, createResp.User.Name, createResp.User.Email, createResp.User.Age)
	}

	// Test 2: Create another user
	fmt.Println("\n2. Creating another user...")
	createResp2, err := client.CreateUser(ctx, &proto.CreateUserRequest{
		Name:     "Jane Smith",
		Email:    "jane@example.com",
		Password: "password123",
		Age:      25,
	})
	if err != nil {
		log.Printf("CreateUser failed: %v", err)
	} else {
		fmt.Printf("   User created: ID=%d, Name=%s, Email=%s, Age=%d\n",
			createResp2.User.Id, createResp2.User.Name, createResp2.User.Email, createResp2.User.Age)
	}

	// Test 3: Get All Users
	fmt.Println("\n3. Getting all users...")
	allUsersResp, err := client.GetAllUsers(ctx, &proto.GetAllUsersRequest{})
	if err != nil {
		log.Printf("GetAllUsers failed: %v", err)
	} else {
		fmt.Printf("   Found %d users:\n", len(allUsersResp.Users))
		for _, user := range allUsersResp.Users {
			fmt.Printf("     - ID: %d, Name: %s, Email: %s, Age: %d\n",
				user.Id, user.Name, user.Email, user.Age)
		}
	}

	// Test 4: Get User by ID
	if createResp != nil && createResp.Success {
		fmt.Printf("\n4. Getting user with ID %d...\n", createResp.User.Id)
		getResp, err := client.GetUser(ctx, &proto.GetUserRequest{Id: createResp.User.Id})
		if err != nil {
			log.Printf("GetUser failed: %v", err)
		} else {
			fmt.Printf("   User found: ID=%d, Name=%s, Email=%s, Age=%d\n",
				getResp.User.Id, getResp.User.Name, getResp.User.Email, getResp.User.Age)
		}

		// Test 5: Update User
		fmt.Printf("\n5. Updating user with ID %d...\n", createResp.User.Id)
		updateResp, err := client.UpdateUser(ctx, &proto.UpdateUserRequest{
			Id:   createResp.User.Id,
			Name: "John Doe Updated",
			Age:  31,
		})
		if err != nil {
			log.Printf("UpdateUser failed: %v", err)
		} else {
			fmt.Printf("   User updated: ID=%d, Name=%s, Email=%s, Age=%d\n",
				updateResp.User.Id, updateResp.User.Name, updateResp.User.Email, updateResp.User.Age)
		}

		// Test 6: Delete User
		fmt.Printf("\n6. Deleting user with ID %d...\n", createResp.User.Id)
		deleteResp, err := client.DeleteUser(ctx, &proto.DeleteUserRequest{Id: createResp.User.Id})
		if err != nil {
			log.Printf("DeleteUser failed: %v", err)
		} else {
			fmt.Printf("   %s\n", deleteResp.Message)
		}
	}

	// Test 7: Get All Users after deletion
	fmt.Println("\n7. Getting all users after deletion...")
	allUsersResp2, err := client.GetAllUsers(ctx, &proto.GetAllUsersRequest{})
	if err != nil {
		log.Printf("GetAllUsers failed: %v", err)
	} else {
		fmt.Printf("   Found %d users:\n", len(allUsersResp2.Users))
		for _, user := range allUsersResp2.Users {
			fmt.Printf("     - ID: %d, Name: %s, Email: %s, Age: %d\n",
				user.Id, user.Name, user.Email, user.Age)
		}
	}

	fmt.Println("\n=== Test completed ===")
}
