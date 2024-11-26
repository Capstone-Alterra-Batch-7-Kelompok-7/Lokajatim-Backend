package response

import "lokajatim/entities"

type RegisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	Email string `json:"email"`
}

func RegisterFromEntities(user entities.User) RegisterResponse {
	return RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
