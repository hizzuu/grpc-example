package repository

import (
	"context"
	"database/sql"

	"github.com/hizzuu/grpc-example-user/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository struct {
	conn sqlHandler
}

func NewUserRepository(sqlHandler sqlHandler) *UserRepository {
	return &UserRepository{
		conn: sqlHandler,
	}
}

func (r *UserRepository) Get(ctx context.Context, id int64) (*domain.User, error) {
	query := `SELECT id, email, name, created_at, updated_at FROM users WHERE id = ?`
	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, id)

	return convertRowToUser(row)
}

func convertRowToUser(row *sql.Row) (*domain.User, error) {
	user := &domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return user, nil
}
