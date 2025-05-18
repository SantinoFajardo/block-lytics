package server

import (
	"context"
	"log"

	userspb "github.com/block-lytics/proto/users/userspb"
	"github.com/block-lytics/users/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersServer struct {
	userspb.UnimplementedUsersServiceServer
	UserRepository db.UserRepository
}

func (s *UsersServer) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.CreateUserResponse, error) {
	log.Println("Creating user with email:", req.Email, "and password:", req.Password)
	id, err := s.UserRepository.CreateUser(req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &userspb.CreateUserResponse{Id: id, Message: "User created successfully!"}, nil
}
