package response

import "lokajatim/entities"

// LoginResponse is the response for the Login controller
// @Description LoginResponse is the response for the Login controller
// @Param ID int true "ID of the user"
// @Param Name string true "Name of the user"
// @Param Email string true "Email of the user"
// @Param Token string true "Token of the user"
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
