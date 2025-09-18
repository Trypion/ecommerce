package repository

import (
	"context"

	"github.com/Trypion/ecommerce/user-service/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}
