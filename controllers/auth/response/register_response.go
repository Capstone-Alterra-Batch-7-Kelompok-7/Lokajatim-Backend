package response

import "lokajatim/entities"

// RegisterResponse is the response for the Register controller
// @Description RegisterResponse is the response for the Register controller
// @Param ID int true "ID of the user"
// @Param Name string true "Name of the user"
// @Param Email string true "Email of the user"
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
