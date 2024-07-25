package ports

import (
	"context"
	"backend/users/internal/core/domain"
)

type UserRepository interface {
	Get(ctx context.Context, id domain.UserID) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}


