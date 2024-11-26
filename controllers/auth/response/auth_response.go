package response

import "lokajatim/entities"

type AuthResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FromEntities(user entities.User) AuthResponse {
	return AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
}
