package request

import "lokajatim/entities"

type UpdateUserRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password,omitempty"`   
	Address     string `json:"address,omitempty"`   
	PhoneNumber string `json:"phone_number,omitempty"` 
	NIK         string `json:"nik,omitempty"`        
}

func (req *UpdateUserRequest) ToEntities() entities.User {
	return entities.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		NIK:         req.NIK,
	}
}
