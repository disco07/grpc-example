package users

import (
	"context"
	"github.com/disco07/gRPC/generated/users"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type Users struct {
	users.UserServiceServer
}

func NewUsers() users.UserServiceServer {
	return &Users{}
}

func (u Users) Create(_ context.Context, req *users.CreateUserRequest) (*emptypb.Empty, error) {
	log.Println("Create user")
	log.Println("Username:", req.GetUsername())
	log.Println("Is active:", req.GetIsActive())
	log.Println("Password:", req.GetPassword())

	return &emptypb.Empty{}, nil
}

func (u Users) Get(_ context.Context, req *users.GetUserRequest) (*users.User, error) {
	if req.GetId() == 1 {
		return &users.User{
			Id:       1,
			Username: "user1",
			IsActive: true,
			Password: "password1",
		}, nil
	}

	return &users.User{}, nil
}
