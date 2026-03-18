package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Kaiman30/AgileWebApp/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO agilewebapp.users (full_name, email)
	VALUES ($1, $2)
	RETURNING id, version, full_name, email;
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.Email)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.Email,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.Email,
	)

	return userDomain, nil
}
