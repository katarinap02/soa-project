package grpc

import (
	"context"
	"database-example/service"
	pb "database-example/proto" // Adjust this import path to match your module
	"log"
	"strings"
)

// StakeholderGrpcServer implements the gRPC service
type StakeholderGrpcServer struct {
	pb.UnimplementedStakeholderServiceServer
	UserService *service.UserService
}

// NewStakeholderGrpcServer creates a new gRPC server instance
func NewStakeholderGrpcServer(userService *service.UserService) *StakeholderGrpcServer {
	return &StakeholderGrpcServer{
		UserService: userService,
	}
}

// GetAllUsers implements the gRPC GetAllUsers method
func (s *StakeholderGrpcServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	log.Println("gRPC GetAllUsers called")

	// Use your existing service logic
	userDTOs, err := s.UserService.GetAllUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return &pb.GetAllUsersResponse{
			Success: false,
			Message: "Failed to get users: " + err.Error(),
			Users:   nil,
		}, nil
	}

	// Convert DTOs to protobuf Users
	var pbUsers []*pb.User
	for _, userDTO := range userDTOs {
		pbUser := &pb.User{
			Id:        userDTO.Id,        // string from your DTO
			Email:     userDTO.Email,     // string from your DTO  
			Name:      userDTO.Username,  // string from your DTO
			Role:      userDTO.Role,      // string from your DTO
		}
		pbUsers = append(pbUsers, pbUser)
	}

	log.Printf("Returning %d users via gRPC", len(pbUsers))
	
	return &pb.GetAllUsersResponse{
		Success: true,
		Message: "Users retrieved successfully via gRPC",
		Users:   pbUsers,
	}, nil
}

// LoginUser implements the gRPC LoginUser method
func (s *StakeholderGrpcServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	log.Printf("gRPC LoginUser called for username: %s", req.Username)

	// Input validation
	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		return &pb.LoginUserResponse{
			User:    nil,
			Token:   "",
		}, nil
	}

	// Use your existing authentication logic - exactly the same as HTTP endpoint
	token, userDTO, err := s.UserService.Authenticate(req.Username, req.Password)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		return &pb.LoginUserResponse{
			User:    nil,
			Token:   "",
		}, nil
	}

	// Convert your DTO to protobuf User
	pbUser := &pb.User{
		Id:        userDTO.Id,        // string from your DTO
		Email:     userDTO.Email,     // string from your DTO
		Name:      userDTO.Username,  // string from your DTO
		Role:      userDTO.Role,      // string from your DTO
	}

	log.Printf("User %s authenticated successfully via gRPC", userDTO.Username)

	return &pb.LoginUserResponse{
		User:    pbUser,
		Token:   token,
	}, nil
}
