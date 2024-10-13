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
	log.Println("Password:", string(req.GetPassword()))
	log.Println("Gender:", req.GetGender())
	log.Println("Emails:", req.GetEmails())

	return &emptypb.Empty{}, nil
}

func (u Users) Get(_ context.Context, req *users.GetUserRequest) (*users.User, error) {
	if req.GetId() == 1 {
		return &users.User{
			Id:       1,
			Username: "john doe",
			IsActive: true,
			Password: "password1",
			Gender:   users.Gender_MALE,
			Emails:   []string{"john.doe@gmail.com"},
		}, nil
	}

	return &users.User{}, nil
}
