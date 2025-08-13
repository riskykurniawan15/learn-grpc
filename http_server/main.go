package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/riskykurniawan15/learn-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HTTPServer struct {
	grpcClient proto.UserServiceClient
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int32  `json:"age"`
}

type UpdateUserRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   int32  `json:"age,omitempty"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewHTTPServer() *HTTPServer {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	return &HTTPServer{
		grpcClient: proto.NewUserServiceClient(conn),
	}
}

func (s *HTTPServer) createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := s.grpcClient.CreateUser(ctx, &proto.CreateUserRequest{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: resp.Success,
		Message: resp.Message,
		Data:    resp.User,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := s.grpcClient.GetUser(ctx, &proto.GetUserRequest{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: resp.Success,
		Message: resp.Message,
		Data:    resp.User,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) getAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := s.grpcClient.GetAllUsers(ctx, &proto.GetAllUsersRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: resp.Success,
		Message: resp.Message,
		Data:    resp.Users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := s.grpcClient.UpdateUser(ctx, &proto.UpdateUserRequest{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: resp.Success,
		Message: resp.Message,
		Data:    resp.User,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := s.grpcClient.DeleteUser(ctx, &proto.DeleteUserRequest{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: resp.Success,
		Message: resp.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *HTTPServer) healthCheck(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "HTTP server is running",
		Data: map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
			"service":   "User Service HTTP Wrapper",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	server := NewHTTPServer()

	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health", server.healthCheck).Methods("GET")

	// User routes
	router.HandleFunc("/users", server.createUser).Methods("POST")
	router.HandleFunc("/users", server.getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", server.getUser).Methods("GET")
	router.HandleFunc("/users/{id}", server.updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", server.deleteUser).Methods("DELETE")

	// CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	port := ":8080"
	fmt.Printf("HTTP server starting on port %s...\n", port)
	fmt.Printf("Health check: http://localhost%s/health\n", port)
	fmt.Printf("API endpoints:\n")
	fmt.Printf("  POST   /users     - Create user\n")
	fmt.Printf("  GET    /users     - Get all users\n")
	fmt.Printf("  GET    /users/{id} - Get user by ID\n")
	fmt.Printf("  PUT    /users/{id} - Update user\n")
	fmt.Printf("  DELETE /users/{id} - Delete user\n")

	log.Fatal(http.ListenAndServe(port, router))
}
