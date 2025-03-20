package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"go-rest-api/internal/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}

type PostgresUserRepository struct {
	Conn *pgx.Conn
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	err := r.Conn.QueryRow(ctx, "SELECT id, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	err := r.Conn.QueryRow(ctx, "INSERT INTO users (email) VALUES ($1) RETURNING id", user.Email).Scan(&user.ID)
	return err
}
