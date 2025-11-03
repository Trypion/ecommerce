package service

import (
	"context"

	"github.com/Trypion/ecommerce/user-service/internal/models"
	"github.com/Trypion/ecommerce/user-service/internal/repository"
	"github.com/Trypion/ecommerce/user-service/internal/utils"
)

type UserService interface {
	Create(ctx context.Context, email string, password string, name string) (*models.User, error)
	Update(ctx context.Context, userID string, email string, name string) (*models.User, error)
	GetById(ctx context.Context, userID string) (*models.User, error)
	Delete(ctx context.Context, userID string) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.AuthLogin, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(ctx context.Context, email string, password string) (*models.AuthLogin, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &models.InvalidCredentialsError{}
	}

	notMatch := utils.CheckPassword(user.Password, password)
	if notMatch {
		return nil, &models.InvalidCredentialsError{}
	}

	signedToken, err := utils.SignJWT(*user)
	if err != nil {
		return nil, err
	}

	return &models.AuthLogin{
		// User:  models.AuthUser{ID: user.ID, Email: user.Email, Name: user.Name, Role: user.Role},
		User:  user,
		Token: signedToken,
	}, nil
}

func (s *userService) Create(ctx context.Context, email string, password string, name string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, &models.UserAlreadyExistsError{Email: email}
	}

	user = &models.User{
		Email:    email,
		Name:     name,
		Password: hashedPassword,
		Role:     "USER",
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(ctx context.Context, userID string, email string, name string) (*models.User, error) {
	user, err := s.repo.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &models.UserNotFoundError{UserID: userID}
	}

	user.Email = email
	user.Name = name

	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetById(ctx context.Context, userID string) (*models.User, error) {
	return s.repo.GetById(ctx, userID)
}

func (s *userService) Delete(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.repo.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &models.UserNotFoundError{UserID: userID}
	}
	if err := s.repo.Delete(ctx, userID); err != nil {
		return nil, err
	}
	return user, nil
}
