package handlers

import (
	"context"
	"time"

	userpb "github.com/Trypion/ecommerce/proto/user"
	"github.com/Trypion/ecommerce/user-service/internal/models"
	"github.com/Trypion/ecommerce/user-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(
	ctx context.Context,
	req *userpb.CreateUserRequest,
) (*userpb.CreateUserResponse, error) {
	user, err := h.service.Create(ctx, req.Email, req.Password, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &userpb.CreateUserResponse{
		User: converUserToProto(user),
	}, nil
}

func converUserToProto(user *models.User) *userpb.User {
	return &userpb.User{
		Id:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

func (h *UserHandler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := h.service.GetById(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	return &userpb.GetUserResponse{
		User: converUserToProto(user),
	}, nil
}
