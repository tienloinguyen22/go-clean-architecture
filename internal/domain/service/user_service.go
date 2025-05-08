package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/entity"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/repository"
)

type IUserService interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
}

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.userRepo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	existingUser, err := s.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to retrieved user for id %s: %w", user.ID, err)
	}
	if existingUser == nil {
		return fmt.Errorf("user not found for id %s", user.ID)
	}

	user.UpdatedAt = time.Now()
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to retrieved user for id %s: %w", id, err)
	}
	if existingUser == nil {
		return fmt.Errorf("user not found for id %s", id)
	}

	return s.userRepo.Delete(ctx, id)
}
