package request

import "lokajatim/entities"

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
