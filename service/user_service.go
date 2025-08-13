package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/riskykurniawan15/learn-grpc/models"
	"github.com/riskykurniawan15/learn-grpc/proto"
	"github.com/riskykurniawan15/learn-grpc/repository"
	"github.com/riskykurniawan15/learn-grpc/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService implements the gRPC UserService interface
type UserService struct {
	proto.UnimplementedUserServiceServer
	userRepo  *repository.UserRepository
	validator *validation.Validator
}

// NewUserService creates a new user service
func NewUserService(validator *validation.Validator) *UserService {
	return &UserService{
		userRepo:  repository.NewUserRepository(),
		validator: validator,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	// Convert proto request to validation struct
	createReq := models.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Age:      int(req.Age),
	}

	// Validate request
	if err := s.validator.ValidateStruct(createReq); err != nil {
		validationErrors := s.validator.GetValidationErrors(err)
		return &proto.CreateUserResponse{
			Success: false,
			Message: "Validation failed: " + strings.Join(validationErrors, "; "),
		}, status.Error(codes.InvalidArgument, "Validation failed")
	}

	// Check if email already exists
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return &proto.CreateUserResponse{
			Success: false,
			Message: "Email already exists",
		}, status.Error(codes.AlreadyExists, "Email already exists")
	}

	// Hash password
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))

	// Create user model
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Age:      int(req.Age),
	}

	// Save to database
	if err := s.userRepo.Create(user); err != nil {
		return &proto.CreateUserResponse{
			Success: false,
			Message: "Failed to create user: " + err.Error(),
		}, status.Error(codes.Internal, "Database error")
	}

	// Convert to proto message
	protoUser := &proto.User{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Age:       int32(user.Age),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return &proto.CreateUserResponse{
		User:    protoUser,
		Message: "User created successfully",
		Success: true,
	}, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	if req.Id <= 0 {
		return &proto.GetUserResponse{
			Success: false,
			Message: "Invalid user ID",
		}, status.Error(codes.InvalidArgument, "Invalid user ID")
	}

	user, err := s.userRepo.GetByID(uint(req.Id))
	if err != nil {
		return &proto.GetUserResponse{
			Success: false,
			Message: "User not found",
		}, status.Error(codes.NotFound, "User not found")
	}

	protoUser := &proto.User{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Age:       int32(user.Age),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return &proto.GetUserResponse{
		User:    protoUser,
		Message: "User retrieved successfully",
		Success: true,
	}, nil
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers(ctx context.Context, req *proto.GetAllUsersRequest) (*proto.GetAllUsersResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return &proto.GetAllUsersResponse{
			Success: false,
			Message: "Failed to retrieve users: " + err.Error(),
		}, status.Error(codes.Internal, "Database error")
	}

	var protoUsers []*proto.User
	for _, user := range users {
		protoUser := &proto.User{
			Id:        int64(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Age:       int32(user.Age),
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		}
		protoUsers = append(protoUsers, protoUser)
	}

	return &proto.GetAllUsersResponse{
		Users:   protoUsers,
		Message: "Users retrieved successfully",
		Success: true,
	}, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	if req.Id <= 0 {
		return &proto.UpdateUserResponse{
			Success: false,
			Message: "Invalid user ID",
		}, status.Error(codes.InvalidArgument, "Invalid user ID")
	}

	// Get existing user
	user, err := s.userRepo.GetByID(uint(req.Id))
	if err != nil {
		return &proto.UpdateUserResponse{
			Success: false,
			Message: "User not found",
		}, status.Error(codes.NotFound, "User not found")
	}

	// Create update request for validation
	updateReq := models.UpdateUserRequest{}

	if req.Name != "" {
		updateReq.Name = req.Name
	}
	if req.Email != "" {
		updateReq.Email = req.Email
	}
	if req.Password != "" {
		updateReq.Password = req.Password
	}
	if req.Age > 0 {
		updateReq.Age = int(req.Age)
	}

	// Validate update request
	if err := s.validator.ValidateStruct(updateReq); err != nil {
		validationErrors := s.validator.GetValidationErrors(err)
		return &proto.UpdateUserResponse{
			Success: false,
			Message: "Validation failed: " + strings.Join(validationErrors, "; "),
		}, status.Error(codes.InvalidArgument, "Validation failed")
	}

	// Check email uniqueness if updating email
	if req.Email != "" && req.Email != user.Email {
		existingUser, _ := s.userRepo.GetByEmail(req.Email)
		if existingUser != nil {
			return &proto.UpdateUserResponse{
				Success: false,
				Message: "Email already exists",
			}, status.Error(codes.AlreadyExists, "Email already exists")
		}
	}

	// Update fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))
	}
	if req.Age > 0 {
		user.Age = int(req.Age)
	}

	// Save changes
	if err := s.userRepo.Update(user); err != nil {
		return &proto.UpdateUserResponse{
			Success: false,
			Message: "Failed to update user: " + err.Error(),
		}, status.Error(codes.Internal, "Database error")
	}

	protoUser := &proto.User{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Age:       int32(user.Age),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	return &proto.UpdateUserResponse{
		User:    protoUser,
		Message: "User updated successfully",
		Success: true,
	}, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	if req.Id <= 0 {
		return &proto.DeleteUserResponse{
			Success: false,
			Message: "Invalid user ID",
		}, status.Error(codes.InvalidArgument, "Invalid user ID")
	}

	// Check if user exists
	_, err := s.userRepo.GetByID(uint(req.Id))
	if err != nil {
		return &proto.DeleteUserResponse{
			Success: false,
			Message: "User not found",
		}, status.Error(codes.NotFound, "User not found")
	}

	// Delete user
	if err := s.userRepo.Delete(uint(req.Id)); err != nil {
		return &proto.DeleteUserResponse{
			Success: false,
			Message: "Failed to delete user: " + err.Error(),
		}, status.Error(codes.Internal, "Database error")
	}

	return &proto.DeleteUserResponse{
		Message: "User deleted successfully",
		Success: true,
	}, nil
}
