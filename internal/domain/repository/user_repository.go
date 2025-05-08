package repository

import (
	"context"

	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/entity"
)

type IUserRepository interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}
