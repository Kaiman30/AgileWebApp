package users_transport_http

import "github.com/Kaiman30/AgileWebApp/internal/core/domain"

type UserDTOResponse struct {
	ID       int    `json:"id"`
	Version  int    `json:"version"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:       user.ID,
		Version:  user.Version,
		FullName: user.FullName,
		Email:    user.Email,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
