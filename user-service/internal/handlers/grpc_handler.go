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

func (h *UserHandler) Login(
	ctx context.Context,
	req *userpb.LoginRequest,
) (*userpb.LoginResponse, error) {
	auth, err := h.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to login user: %v", err)
	}
	return &userpb.LoginResponse{
		User:        converUserToProto(auth.User),
		AccessToken: auth.Token,
	}, nil
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

func (h *UserHandler) GetUser(
	ctx context.Context,
	req *userpb.GetUserRequest,
) (*userpb.GetUserResponse, error) {
	user, err := h.service.GetById(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	return &userpb.GetUserResponse{
		User: converUserToProto(user),
	}, nil
}

func (h *UserHandler) UpdateUser(
	ctx context.Context,
	req *userpb.UpdateUserRequest,
) (*userpb.UpdateUserResponse, error) {
	user, err := h.service.Update(ctx, req.UserId, req.Email, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	return &userpb.UpdateUserResponse{
		User: converUserToProto(user),
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context,
	req *userpb.DeleteUserRequest,
) (*userpb.DeleteUserResponse, error) {
	user, err := h.service.Delete(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	return &userpb.DeleteUserResponse{
		User:    converUserToProto(user),
		Deleted: true,
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
