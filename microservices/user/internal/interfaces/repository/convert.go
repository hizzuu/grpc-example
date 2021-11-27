package repository

import (
	"database/sql"

	"github.com/hizzuu/grpc-example-user/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertRowToUser(row *sql.Row) (*domain.User, error) {
	user := &domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.Email,
		&user.EncryptedPassword,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return user, nil
}
