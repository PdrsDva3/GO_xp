package repository

import (
	"GO_xp/internal/entities"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, user entities.UserCreate) (int, error)
	Get(ctx context.Context, userID int) (*entities.User, error)
	GetPassword(ctx context.Context, login string) (int, string, error)
	UpdatePassword(ctx context.Context, userID int, newPassword string) error
	Delete(ctx context.Context, userID int) error
}
