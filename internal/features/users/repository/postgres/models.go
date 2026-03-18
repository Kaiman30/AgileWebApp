package users_postgres_repository

import "github.com/Kaiman30/AgileWebApp/internal/core/domain"

type UserModel struct {
	ID       int
	Version  int
	FullName string
	Email    string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.Email,
		)
	}

	return userDomains
}
