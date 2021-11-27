package repository

import (
	"context"

	"github.com/hizzuu/grpc-example-user/internal/domain"
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
	query := `SELECT id, email, encrypted_password, name, created_at, updated_at FROM users WHERE id = ?`

	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	return convertRowToUser(row)
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, encrypted_password, name, created_at, updated_at FROM users WHERE email = ?`

	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, email)

	return convertRowToUser(row)
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	query := `INSERT users SET email=?, encrypted_password=?, name=?, created_at=?, updated_at=?`

	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, u.Name, u.Email, u.EncryptedPassword, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id

	return u, nil
}
