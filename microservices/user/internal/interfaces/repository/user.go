package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/hizzuu/grpc-sample-user/internal/domain"
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
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()
	user := &domain.User{}
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			fmt.Println("NOT FOUND")
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}
