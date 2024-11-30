package request

import "lokajatim/entities"

// RegisterRequest is the request for the register endpoint
// @Description RegisterRequest is the request for the register endpoint
// @Param Email string true "Email of the user"
// @Param Password string true "Password of the user"
type RegisterRequest struct {
	Name          string 	`json:"name"`
	Email         string    `json:"email"`
	Password      string 	`json:"password"`
	Role		  string 	`json:"role,omitempty"`
}

func (registerRequest RegisterRequest) ToEntities() entities.User {
	return entities.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		Role:	  registerRequest.Role,
	}
}
