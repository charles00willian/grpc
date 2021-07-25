package services

import (
	"context"
	"fmt"

	"gihub.com/charles00willian/grcp-go-starter/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// Insert database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    req.Id,
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
