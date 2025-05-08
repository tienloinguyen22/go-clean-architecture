package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/entity"
	"github.com/tienloinguyen22/go-clean-architecture/internal/domain/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	user := &entity.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan user row: %w", err)
	}
	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
