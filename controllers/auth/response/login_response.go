package response

import "lokajatim/entities"

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func LoginFromEntities(user entities.User) LoginResponse {
	return LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
}
