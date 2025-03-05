package repository

import (
    "context"
    "go-rest-api/internal/model"
    "github.com/jackc/pgx/v5"
)

type UserRepository interface {
    GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type PostgresUserRepository struct {
    Conn *pgx.Conn
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
    user := &model.User{}
    err := r.Conn.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return user, nil
}
